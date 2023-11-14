package http

import (
	"net/http"

	"github.com/CamiloAvelar/go-user-service/bootstrap"
	"github.com/CamiloAvelar/go-user-service/infrastructure/http/middlewares"
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	"github.com/gorilla/mux"
)

func healthHandler(i infrainterfaces.ServerInjections) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := i.Db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func readinessHandler(i infrainterfaces.ServerInjections) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found\n"))
}

func GetRouter(i *infrainterfaces.ServerInjections) *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.ContentTypeApplicationJsonMiddleware)
	r.Use(middlewares.ErrorMiddleware)

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.HandleFunc("/health", healthHandler(*i))
	r.HandleFunc("/readiness", readinessHandler(*i))

	router := infrainterfaces.Router{
		Injections: i,
		Router:     r,
	}

	bootstrap.SetupRoutes(&router)

	return r
}
