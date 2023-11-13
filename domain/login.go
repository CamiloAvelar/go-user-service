package domain

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type Login struct {
	User
}

type LoginInterface interface {
	Validate() error
	CreateAccessToken(secret string, expiry string) (string, error)
}
