package mongo

import (
	"time"

	"github.com/iAmPlus/microservice/models/apimodels"
)

type Opts struct {
	// time in seconds for considering dialogue records for context.
	RecentDialogueThreshold float64
	// time in seconds for considering dialogue continuation
	NewDialogueThreshold float64
}

// CreateOptsFromConfig creates opts from config
func createOptsFromConfig() Opts {
	return Opts{
		RecentDialogueThreshold: 5 * 60.0,
		NewDialogueThreshold:    3 * 60 * 60.0,
	}
}

// SearchOpts for SearchHistory function.
type SearchOpts struct {
	UserID             string
	Query              string
	Intent             string
	Limit              int
	Offset             int
	IgnorePreformatted bool
	Start              time.Time
	End                time.Time
}

//Manager ... stores/retrieves dialogue records.
type Manager interface {
	Close() error
	Clear() error
	Ping() error
	InserUser(user apimodels.User) error
	InserPost(apimodels.FeedPost) error
	FollowUser(apimodels.FrendRequest) error
	UnFollowUser(apimodels.FrendRequest) error
	GetUserFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.UsersFeed, error)
	GetUserFrendsFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.FriendsFeed, error)
	GetRelatedFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.RelatedFeedResponse, error)
	GetAllUsers() (apimodels.Users, error)
}
