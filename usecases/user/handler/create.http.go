package userhandler

import (
	"encoding/json"
	"net/http"

	"github.com/CamiloAvelar/go-user-service/domain"
)

func (handler handler) CreateHttp(response http.ResponseWriter, request *http.Request) {
	createUserRequest := domain.User{}

	if err := json.NewDecoder(request.Body).Decode(&createUserRequest); err != nil {
		panic(err)
	}

	user, err := handler.usecase.Create(createUserRequest)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(response).Encode(user)
}
