package services

import (
	"github.com/saravana-1010/go-bookstore-users-api/domain/users"
	"github.com/saravana-1010/go-bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if validationErr := user.Validate(); validationErr != nil {
		return nil, validationErr
	}

	if setError := user.Save(); setError != nil {
		return nil, setError
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr){
	// users.
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	getUser := &users.User{
		Id: userId,
	}
	getError := getUser.Get()
	if getError != nil {
		return nil, getError
	}
	return getUser, nil
}