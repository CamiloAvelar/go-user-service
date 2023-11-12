package userhandler

import (
	"encoding/json"
	"net/http"

	userdto "github.com/CamiloAvelar/go-user-service/usecases/user/dto"
)

func (handler handler) Create(response http.ResponseWriter, request *http.Request) {
	createUserRequest := userdto.CreateUserRequest{}

	if err := json.NewDecoder(request.Body).Decode(&createUserRequest); err != nil {
		panic(err)
	}

	user, err := handler.usecase.Create(createUserRequest)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(response).Encode(user)
}
