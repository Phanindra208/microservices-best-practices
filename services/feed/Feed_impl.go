package feedService

import (
	"github.com/iAmPlus/microservice/models/apimodels"
)

//GetUserFeed
func (ls *PaymentService) GetUserFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.UsersFeed, error) {

	return ls.mongo.GetUserFeed(userEmail, Pagestate)

}

//GetUserFrendsFeed
func (ls *PaymentService) GetUserFrendsFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.FriendsFeed, error) {

	return ls.mongo.GetUserFrendsFeed(userEmail, Pagestate)

}

//GetUserFrendsFeed
func (ls *PaymentService) GetRelatedFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.RelatedFeedResponse, error) {

	return ls.mongo.GetRelatedFeed(userEmail, Pagestate)

}

//GetUserFrendsFeed
func (ls *PaymentService) InserPost(post apimodels.FeedPost) error {

	return ls.mongo.InserPost(post)

}
