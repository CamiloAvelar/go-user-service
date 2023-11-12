package userdto

import (
	"github.com/CamiloAvelar/go-user-service/domain/errors"
	"github.com/CamiloAvelar/go-user-service/infrastructure/encryption"
)

type LoginUserDto interface {
	Validate() error
	ComparePassword(password string) bool
	CreateAccessToken(secret string, expiry string) (string, error)
}

type LoginUserRequest struct {
	ID       int64
	Document string `json:"document"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
}

func (l *LoginUserRequest) SetID(id int64) {
	l.ID = id
}

func (l LoginUserRequest) Validate() error {
	if l.Document == "" && l.Email == "" {
		return &errors.ApplicationError{
			Type:       "Validation",
			Message:    "Invalid fields, should have document or email",
			StatusCode: 400,
		}
	}

	if l.Password == "" {
		return &errors.ApplicationError{
			Type:       "Validation",
			Message:    "Password cannot be empty",
			StatusCode: 400,
		}
	}

	return nil
}

func (l LoginUserRequest) ComparePassword(p string) bool {
	password := encryption.Password{
		Password: l.Password,
	}

	err := password.Compare(p)

	return err == nil
}

func (l LoginUserRequest) CreateAccessToken(s string, e int64) (string, error) {
	tokenInformations := encryption.Token{
		ID:     l.ID,
		Secret: s,
		Expiry: e,
	}

	token, err := tokenInformations.CreateAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}
