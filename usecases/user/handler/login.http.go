package userhandler

import (
	"encoding/json"
	"net/http"

	"github.com/CamiloAvelar/go-user-service/domain"
)

func (handler handler) LoginHttp(response http.ResponseWriter, request *http.Request) {
	loginUserRequest := domain.User{}

	if err := json.NewDecoder(request.Body).Decode(&loginUserRequest); err != nil {
		panic(err)
	}

	login, err := handler.usecase.Login(domain.Login{User: loginUserRequest})

	if err != nil {
		panic(err)
	}

	json.NewEncoder(response).Encode(login)
}
