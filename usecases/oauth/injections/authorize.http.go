package oauthinjections

import (
	"net/http"

	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

func AuthorizeHttpHandler(i *infrainterfaces.ServerInjections) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupOauthInjections(i).
			AuthorizeHttp(w, r)
	}
}
