package oauthhandler

import (
	"github.com/CamiloAvelar/go-user-service/domain"
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

type handler struct {
	usecase    domain.OauthUseCase
	injections *infrainterfaces.ServerInjections
}

func New(usecase domain.OauthUseCase, i *infrainterfaces.ServerInjections) domain.OauthHandler {
	return &handler{
		usecase:    usecase,
		injections: i,
	}
}
