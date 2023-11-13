package encryption

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type TokenEncryption interface {
	CreateAccessToken() (string, error)
}

type Token struct {
	ID     int64
	Name   string
	Secret string
	Expiry int64
}

func (t Token) CreateAccessToken() (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(t.Expiry)).Unix()
	claims := &JwtCustomClaims{
		ID:   t.ID,
		Name: t.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(t.Secret))

	if err != nil {
		return "", err
	}

	return s, err
}
