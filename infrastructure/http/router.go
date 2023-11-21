package http

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/CamiloAvelar/go-user-service/bootstrap"
	"github.com/CamiloAvelar/go-user-service/infrastructure/http/middlewares"
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	_ "github.com/go-sql-driver/mysql"
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

func dumpRequest(writer io.Writer, header string, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	writer.Write([]byte("\n" + header + ": \n"))
	writer.Write(data)
	return nil
}

func GetRouter(i *infrainterfaces.ServerInjections) *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.ContentTypeApplicationJsonMiddleware)
	r.Use(middlewares.ErrorMiddleware)

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.HandleFunc("/health", healthHandler(*i))
	r.HandleFunc("/readiness", readinessHandler(*i))

	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		token, err := i.OauthSrv.ValidationBearerToken(r) //TODO: passar para o middleware
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
			"client_id":  token.GetClientID(),
			"user_id":    token.GetUserID(),
		}
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		e.Encode(data)
	})

	router := infrainterfaces.Router{
		Injections: i,
		Router:     r,
	}

	bootstrap.SetupRoutes(&router)

	return r
}
