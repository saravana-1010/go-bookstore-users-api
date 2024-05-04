package users

import (
	"strings"

	"github.com/saravana-1010/go-bookstore-users-api/utils/errors"
)

type User struct {
	Id 					int64		`json:"id"`
	FirstName 			string		`json:"firstName"`
	LastName 			string		`json:"lastName"`
	Email 				string		`json:"email"`
	DateCreated 		string		`json:"dateCreated"`
}

func (user *User) Validate() *errors.RestErr{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}