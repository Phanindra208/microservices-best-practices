package userService

import "github.com/iAmPlus/microservice/models/apimodels"

type Mongo interface {
	InserUser(user apimodels.User) error

	FollowUser(apimodels.FrendRequest) error
	UnFollowUser(apimodels.FrendRequest) error
	GetAllUsers() (apimodels.Users, error)
}

type PaymentService struct {
	mongo Mongo
}

// New creates new Feed Service.
func New(m Mongo) *PaymentService {
	return &PaymentService{m}
}
