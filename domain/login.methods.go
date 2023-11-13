package domain

import (
	"github.com/CamiloAvelar/go-user-service/domain/errors"
	"github.com/CamiloAvelar/go-user-service/infrastructure/encryption"
)

func (l Login) Validate() error {
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

func (l Login) CreateAccessToken(s string, e int64) (string, error) {
	tokenInformations := encryption.Token{
		ID:     l.ID,
		Name:   l.Name,
		Secret: s,
		Expiry: e,
	}

	token, err := tokenInformations.CreateAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}
