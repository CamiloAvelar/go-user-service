package userinjections

import (
	"net/http"

	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

func CreateHandler(i infrainterfaces.HttpServerInjections) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupUserInjections(i).
			Create(w, r)
	}
}

func LoginHandler(i infrainterfaces.HttpServerInjections) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupUserInjections(i).
			Login(w, r)
	}
}
