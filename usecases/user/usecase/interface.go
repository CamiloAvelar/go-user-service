package userusecase

import (
	"github.com/CamiloAvelar/go-user-service/domain"
	"github.com/CamiloAvelar/go-user-service/infrastructure/config"
)

type usecase struct {
	repository domain.UserRepository
	config     config.Config
}

func New(repository domain.UserRepository, config config.Config) domain.UserUseCase {
	return &usecase{
		repository: repository,
		config:     config,
	}
}
