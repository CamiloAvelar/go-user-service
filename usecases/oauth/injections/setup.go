package oauthinjections

import (
	"github.com/CamiloAvelar/go-user-service/domain"
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	oauthhandler "github.com/CamiloAvelar/go-user-service/usecases/oauth/handler"
	oauthusecase "github.com/CamiloAvelar/go-user-service/usecases/oauth/usecase"
	userrepository "github.com/CamiloAvelar/go-user-service/usecases/user/repository"
)

func setupOauthInjections(i *infrainterfaces.ServerInjections) domain.OauthHandler {
	userRepository := userrepository.New(i.Db)
	oauthUsecase := oauthusecase.New(userRepository, i.Config, i.User)
	return oauthhandler.New(oauthUsecase, i)
}
