package services

import (
	"github.com/thirumalaivasan95/bookstore_users-api/domain/users"
	"github.com/thirumalaivasan95/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// return &user, nil
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
	// return &user, &errors.RestErr{
	// 	Status: http.StatusInternalServerError,
	// }
}
