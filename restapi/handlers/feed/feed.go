package feedhandler

import (
	"github.com/iAmPlus/microservice/models/apimodels"
)

// PaymentService Process Payment.
type FeedService interface {
	GetUserFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.UsersFeed, error)
	GetUserFrendsFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.FriendsFeed, error)
	GetRelatedFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.RelatedFeedResponse, error)
	InserPost(apimodels.FeedPost) error
}

var feedService FeedService

// Init initializes package.
func Init(ls FeedService) {
	feedService = ls
}
