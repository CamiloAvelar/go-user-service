package userusecase

import (
	"github.com/CamiloAvelar/go-user-service/domain"
	"github.com/CamiloAvelar/go-user-service/infrastructure/config"
)

type usecase struct {
	repository domain.UserRepository
	config     config.Config
	user       domain.User
}

func New(repository domain.UserRepository, config config.Config, user domain.User) domain.UserUseCase {
	return &usecase{
		repository: repository,
		config:     config,
		user:       user,
	}
}
