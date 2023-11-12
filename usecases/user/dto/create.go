package userdto

import (
	"net/mail"

	"github.com/CamiloAvelar/go-user-service/domain/errors"
	"github.com/CamiloAvelar/go-user-service/infrastructure/encryption"
	"github.com/klassmann/cpfcnpj"
)

type CreateUserDto interface {
	Validate() error
	EncryptPassword() error
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Document string `json:"document"`
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

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return &errors.ApplicationError{
			Type:       "Validation",
			Message:    "Invalid email",
			StatusCode: 400,
		}
	}

	if !(cpfcnpj.ValidateCNPJ(u.Document) || cpfcnpj.ValidateCPF(u.Document)) {
		return &errors.ApplicationError{
			Type:       "Validation",
			Message:    "Invalid document",
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

	if len(u.Password) < 8 {
		return &errors.ApplicationError{
			Type:       "Validation",
			Message:    "Insecure password",
			StatusCode: 400,
		}
	}

	return nil
}

func (u *CreateUserRequest) EncryptPassword() error {
	password := encryption.Password{
		Password: u.Password,
	}

	encrypted, err := password.Encrypt()

	if err != nil {
		return err
	}

	u.Password = encrypted

	return nil
}
