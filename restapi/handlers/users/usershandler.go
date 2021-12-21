package usershandlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iAmPlus/microservice/restapi/operations/users"
	"github.com/iAmPlus/microservice/restapi/responder"
)

// InserUser.
func InserUser(params users.CreateUserParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	userService.InserUser(*params.Payload)

	return resp.OK("Inserted User")

}

// FollowUser.
func FollowUser(params users.FollowUserParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	userService.FollowUser(*params.Payload)

	return resp.OK("Sucessfully Follow User")

}

// UnFollowUser.
func UnFollowUser(params users.UnFollowUserParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	userService.UnFollowUser(*params.Payload)
	return resp.OK("Sucessfully UnFollow User")

}

// GetAllUsers.
func GetAllUsers(params users.GetUsersParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	result, err := userService.GetAllUsers()
	if err != nil {
		return resp.Status(500).Error(500, err.Error())
	}
	return resp.OK(result)

}
