package encryption

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type TokenEncryption interface {
	CreateAccessToken() (string, error)
	ValidateToken() error
}

type Token struct {
	ID     int64
	Name   string
	Hash   string
	Secret string
	Expiry int64
}

type AccessToken struct {
	Hash      string
	ExpiresAt int64
}

func (t Token) CreateAccessToken() (AccessToken, error) {
	exp := time.Now().Add(time.Hour * time.Duration(t.Expiry))
	claims := &JwtCustomClaims{
		ID:   strconv.Itoa(int(t.ID)),
		Name: t.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(t.Secret))

	if err != nil {
		return AccessToken{}, err
	}

	return AccessToken{Hash: s, ExpiresAt: exp.Unix()}, err
}

func (t *Token) ValidateToken() error {
	token, err := jwt.Parse(t.Hash, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.Secret), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return fmt.Errorf("invalid token")
	}

	id, _ := strconv.Atoi(claims["id"].(string))

	t.ID = int64(id)
	t.Name = claims["name"].(string)

	return nil
}
