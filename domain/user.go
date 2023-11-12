package domain

import (
	"net/http"

	userdto "github.com/CamiloAvelar/go-user-service/usecases/user/dto"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(userRequest userdto.CreateUserRequest) (int64, error)
	FindByEmailOrDocument(email, document string) (User, error)
}

type UserUseCase interface {
	Create(userRequest userdto.CreateUserRequest) (int64, error)
	Login(loginRequest userdto.LoginUserRequest) (userdto.LoginUserResponse, error)
}

type UserHandler interface {
	Create(response http.ResponseWriter, request *http.Request)
	Login(response http.ResponseWriter, request *http.Request)
}
