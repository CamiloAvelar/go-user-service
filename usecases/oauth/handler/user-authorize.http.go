package oauthhandler

import (
	"fmt"
	"net/http"

	"github.com/go-session/session"
)

func (handler handler) UserAuthorizeHttp(response http.ResponseWriter, request *http.Request) (userID string, err error) {
	store, err := session.Start(request.Context(), response, request)
	if err != nil {
		return "", err
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if request.Form == nil {
			request.ParseForm()
		}

		store.Set("ReturnUri", request.Form)
		store.Save()

		response.Header().Set("Location", "/oauth/login")
		response.WriteHeader(http.StatusFound)
		return "", nil
	}

	store.Delete("LoggedInUserID") //TODO: why delete? undestand
	store.Save()

	userID = fmt.Sprintf("%.f", uid.(float64))

	return userID, nil
}
