package mongo

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/iAmPlus/microservice/models/apimodels"
	"github.com/iAmPlus/microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	errMissingDBCredentials = fmt.Errorf("Database Credentials Missing")
	maxDuration             = 100 * time.Second
	indexDuration           = 4 * time.Minute
)

// NewnewsManager creates a user manager that stores user data in memory.
func NewnewsManager(DBUrl string, DBname string) (Manager, error) {
	dbOpts := CreateOptsFromConfig(DBUrl, DBname)
	if dbOpts.DatabaseURI == "" || dbOpts.DatabaseName == "" {
		return nil, errMissingDBCredentials
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(dbOpts.DatabaseURI))
	if err != nil {
		return nil, fmt.Errorf("Unable to create MONGO DB client, URL %s, opts : %v : %v", dbOpts.DatabaseURI, dbOpts, err)
	}
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("Unable to create MONGO DB connection, URL %s, opts : %v : %v", dbOpts.DatabaseURI, dbOpts, err)
	}

	um := &manager{dbOpts: dbOpts, client: client}
	err = um.loadTables()
	if err != nil {
		return nil, fmt.Errorf("Unable to create MONGO DB user manager URL %s, opts : %v : %v", dbOpts.DatabaseURI, dbOpts, err)
	}
	return um, nil
}

type manager struct {
	dbOpts        DatabaseOpts
	client        *mongo.Client
	users         *mongo.Collection
	posts         *mongo.Collection
	related_posts *mongo.Collection
}

//DatabaseOpts options for passing database details
type DatabaseOpts struct {
	// Uri of databse to connect to
	DatabaseURI string
	// Name of database to connect
	DatabaseName      string
	Users_coll        string
	Posts_coll        string
	Related_post_coll string
}

// CreateOptsFromConfig creates opts from config
func CreateOptsFromConfig(DBUrl string, DBname string) DatabaseOpts {
	return DatabaseOpts{
		DatabaseURI:       DBUrl,
		DatabaseName:      DBname,
		Users_coll:        "users",
		Posts_coll:        "posts",
		Related_post_coll: "related",
	}
}

func createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), maxDuration)
}

func (u *manager) loadTables() error {

	ctx, cancelFunc := context.WithTimeout(context.Background(), indexDuration)
	defer cancelFunc()
	var err error
	database := u.client.Database(u.dbOpts.DatabaseName)

	u.posts = database.Collection(u.dbOpts.Posts_coll)
	_, err = u.posts.Indexes().CreateMany(ctx, DBIndex())
	u.users = database.Collection(u.dbOpts.Users_coll)
	_, err = u.posts.Indexes().CreateMany(ctx, DBIndex())
	u.related_posts = database.Collection(u.dbOpts.Related_post_coll)
	_, err = u.posts.Indexes().CreateMany(ctx, DBIndex())

	if err != nil {
		return fmt.Errorf("Unable to create index on dialog table : %v", err)
	}

	return nil
}

func (u *manager) Close() error {
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	return u.client.Disconnect(ctx)
}

func (u *manager) Clear() error {
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	err := u.posts.Drop(ctx)
	if err != nil {
		return fmt.Errorf("Unable to drop feed collection : %v", err)
	}

	err = u.loadTables()
	return fmt.Errorf("Unable to recreate tables : %v", err)
}

func (u *manager) Ping() error {
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	return u.client.Ping(ctx, nil)
}

func DBIndex() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.M{"actor": 1},
			Options: &options.IndexOptions{
				Name: utils.StringPtr("actor"),
				Collation: &options.Collation{
					Locale:   "en",
					Strength: 2,
				},
			},
		},
	}
}

func (u *manager) InserUser(user apimodels.User) error {
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	var collection *mongo.Collection
	collection = u.users
	_, err := collection.InsertOne(ctx, user)
	return err
}
func (u *manager) InserPost(feed apimodels.FeedPost) error {
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	var collection *mongo.Collection
	collection = u.posts
	_, err := collection.InsertOne(ctx, feed)
	if feed.Verb == "post" {
		relatedFeed := apimodels.RelatedFeed{}
		relatedFeed.Actor = feed.Actor
		relatedFeed.Object = feed.Object
		relatedFeed.Verb = feed.Verb
		_, err := u.related_posts.InsertOne(ctx, relatedFeed)
		return err
	} else {
		relatedFeed := apimodels.RelatedFeed{}
		u.related_posts.FindOne(ctx, bson.M{"object": feed.Object}).Decode(&relatedFeed)
		if relatedFeed.Actor != "" {
			relatedFeed.Related = append(relatedFeed.Related, &feed)
			result, err := u.related_posts.UpdateOne(
				ctx,
				bson.M{"object": feed.Object},
				bson.D{
					{"$set", bson.D{{"related", relatedFeed.Related}}},
				},
			)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
		}
	}
	return err
}
func (u *manager) FollowUser(fr apimodels.FrendRequest) error {
	user := apimodels.User{}
	query := bson.M{"useremailid": fr.SourceUserEmailID}
	user_collection := u.users
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	user_collection.FindOne(ctx, query).Decode(&user)
	if user.UserEmailID != "" {
		user.Friends = append(user.Friends, fr.TargetUserEmailID)
		result, err := user_collection.UpdateOne(
			ctx,
			bson.M{"useremailid": fr.SourceUserEmailID},
			bson.D{
				{"$set", bson.D{{"friends", user.Friends}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	} else {
		return fmt.Errorf("user not found")
	}
	return nil
}
func (u *manager) UnFollowUser(fr apimodels.FrendRequest) error {
	user := apimodels.User{}
	query := bson.M{"useremailid": fr.SourceUserEmailID}
	user_collection := u.users
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	user_collection.FindOne(ctx, query).Decode(&user)
	var userar []string
	if user.UserEmailID != "" {
		for _, v := range user.Friends {
			if v != fr.TargetUserEmailID {
				userar = append(userar, v)
				break
			}
		}
		result, err := user_collection.UpdateOne(
			ctx,
			bson.M{"useremailid": fr.SourceUserEmailID},
			bson.D{
				{"$set", bson.D{{"friends", userar}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	} else {
		return fmt.Errorf("user not found")
	}
	return nil
}
func ToIntf(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	intf := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		intf[i] = v.Index(i).Interface()
	}
	return intf
}
func (u *manager) GetUserFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.UsersFeed, error) {

	var result []*apimodels.FeedPost
	pageState := apimodels.PaginationData{}
	query := bson.M{"actor": userEmail}
	usersFeed := apimodels.UsersFeed{}
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	collection := u.posts
	var limit = int64(Pagestate.PageSize)
	var page = int64(Pagestate.PageID)
	var skip int64

	if page > 0 {
		skip = (page - 1) * limit
	} else {
		skip = page
	}
	col := options.Collation{}
	col.Locale = "en"
	col.Strength = 2
	findopts := &options.FindOptions{
		Collation: &col,
		Limit:     &limit,
		Skip:      &skip,
		Sort:      bson.D{{"createdat", -1}},
	}
	cursor, err := collection.Find(ctx, query, findopts)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &result)
	if err != nil {
		return usersFeed, err
	}
	total, _ := collection.CountDocuments(context.Background(), query)
	count := int64(total)
	usersFeed.MyFeed = result
	pageState.Page = page
	pageState.Total = count
	usersFeed.PageState = &pageState
	return usersFeed, nil

}
func (u *manager) GetUserFrendsFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.FriendsFeed, error) {

	var result []*apimodels.FeedPost
	pageState := apimodels.PaginationData{}
	query := bson.M{"useremailid": userEmail}
	usersFeed := apimodels.FriendsFeed{}
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	collection := u.posts

	var limit = int64(Pagestate.PageSize)
	var page = int64(Pagestate.PageID)
	var skip int64

	if page > 0 {
		skip = (page - 1) * limit
	} else {
		skip = page
	}
	col := options.Collation{}
	col.Locale = "en"
	col.Strength = 2
	findopts := &options.FindOptions{
		Collation: &col,
		Limit:     &limit,
		Skip:      &skip,
		Sort:      bson.D{{"createdat", -1}},
	}
	user := apimodels.User{}

	user_collection := u.users

	user_collection.FindOne(ctx, query).Decode(&user)
	if user.Friends != nil {
		query_friends := bson.M{}
		query_friends["actor"] = bson.D{{"$in", user.Friends}}
		cursor, err := collection.Find(ctx, query_friends, findopts)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(ctx)

		err = cursor.All(ctx, &result)
		if err != nil {
			return usersFeed, err
		}
		total, _ := collection.CountDocuments(context.Background(), query_friends)
		count := int64(total)
		usersFeed.FriendsFeed = result
		pageState.Page = page
		pageState.Total = count
		usersFeed.PageState = &pageState
	}
	return usersFeed, nil

}
func (u *manager) GetRelatedFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.RelatedFeedResponse, error) {
	var result []*apimodels.RelatedFeed
	pageState := apimodels.PaginationData{}
	query := bson.M{"actor": userEmail}
	usersFeed := apimodels.RelatedFeedResponse{}
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	collection := u.related_posts
	var limit = int64(Pagestate.PageSize)
	var page = int64(Pagestate.PageID)
	var skip int64
	if page > 0 {
		skip = (page - 1) * limit
	} else {
		skip = page
	}
	col := options.Collation{}
	col.Locale = "en"
	col.Strength = 2
	findopts := &options.FindOptions{
		Collation: &col,
		Limit:     &limit,
		Skip:      &skip,
		Sort:      bson.D{{"createdat", -1}},
	}
	cursor, err := collection.Find(ctx, query, findopts)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &result)
	if err != nil {
		return usersFeed, err
	}
	total, _ := collection.CountDocuments(context.Background(), query)
	count := int64(total)
	usersFeed.RelatedFeed = result
	pageState.Page = page
	pageState.Total = count
	usersFeed.PageState = &pageState
	return usersFeed, nil
}
func (u *manager) GetAllUsers() (apimodels.Users, error) {

	var decode []*apimodels.User
	var result apimodels.Users
	ctx, cancelFunc := createContext()
	defer cancelFunc()
	col := options.Collation{}
	col.Locale = "en"
	col.Strength = 2
	findopts := &options.FindOptions{
		Collation: &col,
	}
	collection := u.users
	cursor, err := collection.Find(ctx, bson.D{}, findopts)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &decode)
	result.User = decode
	return result, nil

}

// GetDateTime Get time from smftdate
func getDateTime(smftdatetime string) time.Time {
	time, _ := time.Parse("2006-01-02T15:04:05.000Z07:00", smftdatetime)

	return time
}
