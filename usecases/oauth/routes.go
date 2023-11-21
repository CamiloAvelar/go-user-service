package oauth

import (
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	oauthinjections "github.com/CamiloAvelar/go-user-service/usecases/oauth/injections"
)

func SetupRoutes(i *infrainterfaces.Router) {
	publicRoutes := i.Router.PathPrefix("/oauth").Subrouter()

	i.Injections.OauthSrv.SetUserAuthorizationHandler(
		oauthinjections.UserAuthorizeHttpHandler(i.Injections),
	)

	publicRoutes.HandleFunc("/login",
		oauthinjections.LoginHttpHandler(i.Injections),
	).Methods("GET", "POST")

	publicRoutes.HandleFunc("/consent",
		oauthinjections.ConsentHttpHandler(i.Injections),
	).Methods("GET", "POST")

	publicRoutes.HandleFunc("/authorize",
		oauthinjections.AuthorizeHttpHandler(i.Injections),
	).Methods("GET", "POST")

	publicRoutes.HandleFunc("/token",
		oauthinjections.TokenHttpHandler(i.Injections),
	).Methods("GET", "POST")
}
