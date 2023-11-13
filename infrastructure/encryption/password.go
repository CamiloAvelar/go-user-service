package encryption

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryption interface {
	Encrypt() (string, error)
	Compare(hashed string) error
}

type Password struct {
	Password string
}

func (p Password) Encrypt() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), 10)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (p Password) Compare(hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p.Password))
}
