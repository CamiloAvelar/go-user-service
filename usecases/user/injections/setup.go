package userinjections

import (
	"github.com/CamiloAvelar/go-user-service/domain"
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	userhandler "github.com/CamiloAvelar/go-user-service/usecases/user/handler"
	userrepository "github.com/CamiloAvelar/go-user-service/usecases/user/repository"
	userusecase "github.com/CamiloAvelar/go-user-service/usecases/user/usecase"
)

func setupUserInjections(i infrainterfaces.ServerInjections) domain.UserHandler {
	userRepository := userrepository.New(i.Db)
	userUsecase := userusecase.New(userRepository, i.Config)
	return userhandler.New(userUsecase)
}
