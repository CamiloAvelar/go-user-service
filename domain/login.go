package domain

type LoginResponse struct {
	TokenType        string `json:"token_type"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	AccessExpiresAt  int64  `json:"access_token_expires_at"`
	RefreshExpiresAt int64  `json:"refresh_token_expires_at"`
}

type Login struct {
	User
}

type LoginInterface interface {
	Validate() error
	CreateAccessToken(secret string, expiry string) (string, error)
}
