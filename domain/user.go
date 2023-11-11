package domain

import (
	"net/http"

	userdto "github.com/CamiloAvelar/go-user-service/usecases/user/dto"
)

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(userRequest userdto.CreateUserRequest) (int64, error)
}

type UserUseCase interface {
	Create(userRequest userdto.CreateUserRequest) (int64, error)
}

type UserHandler interface {
	Create(response http.ResponseWriter, request *http.Request)
}
