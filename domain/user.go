package domain

import (
	"net/http"
)

type UserInterface interface {
	Validate() error
	ValidatePersisted() error
	EncryptPassword() error
	ComparePassword(p string) bool
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(userRequest User) (int64, error)
	FindByEmailOrDocument(email, document string) (User, error)
	FindByID(id int64) (User, error)
}

type UserUseCase interface {
	Create(userRequest User) (int64, error)
	Login(loginRequest Login) (*LoginResponse, error)
}

type UserHandler interface {
	CreateHttp(response http.ResponseWriter, request *http.Request)
	LoginHttp(response http.ResponseWriter, request *http.Request)
}
