package user

import (
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	userinjections "github.com/CamiloAvelar/go-user-service/usecases/user/injections"
)

func SetupRoutes(i *infrainterfaces.Router) {
	publicRoutes := i.Router.PathPrefix("/user").Subrouter()

	publicRoutes.HandleFunc("/create",
		userinjections.CreateHttpHandler(i.Injections),
	).Methods("POST")

	publicRoutes.HandleFunc("/login",
		userinjections.LoginHttpHandler(i.Injections),
	).Methods("POST")
}
