package user

import (
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	userinjections "github.com/CamiloAvelar/go-user-service/usecases/user/injections"
)

func SetupRoutes(i infrainterfaces.Router) {
	routes := i.Router.PathPrefix("/user").Subrouter()

	routes.HandleFunc("/create",
		userinjections.CreateHttpHandler(i.Injections),
	).Methods("POST")

	routes.HandleFunc("/login",
		userinjections.LoginHttpHandler(i.Injections),
	).Methods("POST")
}
