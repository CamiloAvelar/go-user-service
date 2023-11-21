package domain

import "net/http"

type LoginResponse struct {
	User User
}

type Login struct {
	User
}

type LoginInterface interface {
	Validate() error
	// CreateAccessToken(secret string, expiry string) (string, error)
}

type OauthUseCase interface {
	Login(loginRequest Login) (*LoginResponse, error)
	Consent() (bool, error)
}

type OauthHandler interface {
	LoginHttp(response http.ResponseWriter, request *http.Request)
	ConsentHttp(response http.ResponseWriter, request *http.Request)
	AuthorizeHttp(response http.ResponseWriter, request *http.Request)
	UserAuthorizeHttp(response http.ResponseWriter, request *http.Request) (userID string, err error)
	TokenHttp(response http.ResponseWriter, request *http.Request)
}
