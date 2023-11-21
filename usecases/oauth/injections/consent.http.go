package oauthinjections

import (
	"net/http"

	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

func ConsentHttpHandler(i *infrainterfaces.ServerInjections) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupOauthInjections(i).
			ConsentHttp(w, r)
	}
}
