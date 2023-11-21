package oauthhandler

import (
	"net/http"

	"github.com/go-session/session"
)

func (handler handler) ConsentHttp(response http.ResponseWriter, request *http.Request) {
	store, err := session.Start(request.Context(), response, request)
	if err != nil {
		panic(err)
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		response.Header().Set("Location", "/oauth/login")
		response.WriteHeader(http.StatusFound)
		return
	}

	outputHTML(response, request, "static/consent.html")
}
