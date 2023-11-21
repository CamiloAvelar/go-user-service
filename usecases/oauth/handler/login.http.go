package oauthhandler

import (
	"net/http"

	"github.com/CamiloAvelar/go-user-service/domain"
	"github.com/go-session/session"
)

func (handler handler) LoginHttp(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		outputHTML(response, request, "static/login.html")
		return
	}

	store, err := session.Start(request.Context(), response, request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	loginUserRequest := domain.User{}

	if err := request.ParseForm(); err != nil {
		panic(err)
	}

	loginUserRequest.Email = request.Form.Get("email")
	loginUserRequest.Password = request.Form.Get("password")

	login, err := handler.usecase.Login(domain.Login{User: loginUserRequest})

	if err != nil {
		panic(err)
	}

	store.Set("LoggedInUserID", login.User.ID)
	store.Save()

	response.Header().Set("Location", "/oauth/consent")
	response.WriteHeader(http.StatusFound)
}
