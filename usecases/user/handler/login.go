package userhandler

import (
	"encoding/json"
	"net/http"

	userdto "github.com/CamiloAvelar/go-user-service/usecases/user/dto"
)

func (handler handler) Login(response http.ResponseWriter, request *http.Request) {
	loginUserRequest := userdto.LoginUserRequest{}

	if err := json.NewDecoder(request.Body).Decode(&loginUserRequest); err != nil {
		panic(err)
	}

	login, err := handler.usecase.Login(loginUserRequest)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(response).Encode(login)
}
