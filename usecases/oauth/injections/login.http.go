package oauthinjections

import (
	"net/http"

	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

func LoginHttpHandler(i *infrainterfaces.ServerInjections) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupOauthInjections(i).
			LoginHttp(w, r)
	}
}
