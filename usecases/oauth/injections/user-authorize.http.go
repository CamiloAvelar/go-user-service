package oauthinjections

import (
	"net/http"

	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

func UserAuthorizeHttpHandler(i *infrainterfaces.ServerInjections) func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	return func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		return setupOauthInjections(i).
			UserAuthorizeHttp(w, r)
	}
}
