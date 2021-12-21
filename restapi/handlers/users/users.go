package usershandlers

import (
	"github.com/iAmPlus/microservice/models/apimodels"
)

// PaymentService Process Payment.
type UserService interface {
	InserUser(user apimodels.User) error
	FollowUser(apimodels.FrendRequest) error
	UnFollowUser(apimodels.FrendRequest) error
	GetAllUsers() (apimodels.Users, error)
}

var userService UserService

// Init initializes package.
func Init(ls UserService) {
	userService = ls
}
