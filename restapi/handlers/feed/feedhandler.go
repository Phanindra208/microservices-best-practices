package feedhandler

import (
	"github.com/iAmPlus/microservice/restapi/operations/feed"
	"github.com/iAmPlus/microservice/restapi/responder"

	"github.com/go-openapi/runtime/middleware"
)

//GetUserFrendsFeed
func GetUserFeed(params feed.UserFeedParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	result, err := feedService.GetUserFeed(*params.UserEmailID, *params.Payload)
	if err != nil {
		return resp.Status(500).Error(500, err.Error())
	}
	return resp.OK(result)

}

// GetUserFrendsFeed
func GetUserFrendsFeed(params feed.FriendsFeedParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	result, err := feedService.GetUserFrendsFeed(*params.UserEmailID, *params.Payload)
	if err != nil {
		return resp.Status(500).Error(500, err.Error())
	}
	return resp.OK(result)

}

// GetRelatedFeed
func GetRelatedFeed(params feed.GetRelatedFeedParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	result, err := feedService.GetRelatedFeed(*params.UserEmailID, *params.Payload)
	if err != nil {
		return resp.Status(500).Error(500, err.Error())
	}
	return resp.OK(result)

}

// InserPost
func InserPost(params feed.CreateFeedParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	feedService.InserPost(*params.Payload)

	return resp.OK("Post Created")

}
