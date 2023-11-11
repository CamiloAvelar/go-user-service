package http

import (
	"net/http"

	"github.com/CamiloAvelar/go-user-service/bootstrap"
	"github.com/CamiloAvelar/go-user-service/infrastructure/http/middlewares"
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found\n"))
}

func GetRouter(i infrainterfaces.HttpServerInjections) *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.ContentTypeApplicationJsonMiddleware)
	r.Use(middlewares.ErrorMiddleware)

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/readiness", readinessHandler)

	router := infrainterfaces.Router{
		Injections: i,
		Router:     r,
	}

	bootstrap.SetupRoutes(router)

	return r
}
