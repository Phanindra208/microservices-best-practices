package userService

import (
	"github.com/iAmPlus/microservice/models/apimodels"
)

//GetUserFeed
func (ls *PaymentService) InserUser(user apimodels.User) error {

	return ls.mongo.InserUser(user)

}

func (ls *PaymentService) FollowUser(fs apimodels.FrendRequest) error {

	return ls.mongo.FollowUser(fs)

}
func (ls *PaymentService) UnFollowUser(fs apimodels.FrendRequest) error {

	return ls.mongo.UnFollowUser(fs)

}
func (ls *PaymentService) GetAllUsers() (apimodels.Users, error) {

	return ls.mongo.GetAllUsers()

}
