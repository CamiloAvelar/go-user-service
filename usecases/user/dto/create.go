package userdto

import (
	"github.com/CamiloAvelar/go-user-service/domain/errors"
)

type ValidUser interface {
	Validate() error
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u CreateUserRequest) Validate() error {
	if u.Email == "" {
		return &errors.ApplicationError{
			Type:       "Validation",
			Message:    "Email cannot be empty",
			StatusCode: 400,
		}
	}

	if u.Password == "" {
		return &errors.ApplicationError{
			Type:       "Validation",
			Message:    "Password cannot be empty",
			StatusCode: 400,
		}
	}

	return nil
}
