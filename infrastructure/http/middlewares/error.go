package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/CamiloAvelar/go-user-service/domain/errors"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case *errors.ApplicationError:
					httpError := errors.HttpError{
						Error: errors.HttpErrorObject{
							Type:    e.Type,
							Message: e.Message,
						},
					}
					w.WriteHeader(e.StatusCode)
					json.NewEncoder(w).Encode(httpError)
				case error:
					httpError := errors.HttpError{
						Error: errors.HttpErrorObject{
							Type:    "Unexpected",
							Message: e.Error(),
						},
					}
					w.WriteHeader(500)
					json.NewEncoder(w).Encode(httpError)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
