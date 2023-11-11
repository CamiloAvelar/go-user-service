package userhandler

import "github.com/CamiloAvelar/go-user-service/domain"

type handler struct {
	usecase domain.UserUseCase
}

func New(usecase domain.UserUseCase) domain.UserHandler {
	return &handler{
		usecase: usecase,
	}
}
