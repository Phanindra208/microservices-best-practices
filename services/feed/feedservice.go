package feedService

import "github.com/iAmPlus/microservice/models/apimodels"

type Mongo interface {
	GetUserFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.UsersFeed, error)
	GetUserFrendsFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.FriendsFeed, error)
	GetRelatedFeed(userEmail string, Pagestate apimodels.PageState) (apimodels.RelatedFeedResponse, error)
	InserPost(apimodels.FeedPost) error
}

type PaymentService struct {
	mongo Mongo
}

// New creates new Feed Service.
func New(m Mongo) *PaymentService {
	return &PaymentService{m}
}
