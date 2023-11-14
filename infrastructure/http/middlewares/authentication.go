package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/CamiloAvelar/go-user-service/domain"
	"github.com/CamiloAvelar/go-user-service/domain/errors"
	"github.com/CamiloAvelar/go-user-service/infrastructure/encryption"
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	userrepository "github.com/CamiloAvelar/go-user-service/usecases/user/repository"
)

func AuthenticationMiddleware(i *infrainterfaces.ServerInjections) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			t := strings.Split(authHeader, " ")

			if len(t) != 2 || t[0] != "Bearer" {
				log.Println("Wrong Authorization header")
				httpError := errors.HttpError{
					Error: errors.HttpErrorObject{
						Type:    "Authentication",
						Message: "Unauthorized",
					},
				}

				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(httpError)
				return
			}

			authToken := t[1]

			token := encryption.Token{Hash: authToken, Secret: i.Config.AccessTokenSecret}

			if err := token.ValidateToken(); err != nil {
				log.Printf("Error validating token: %v\n", err.Error())
				httpError := errors.HttpError{
					Error: errors.HttpErrorObject{
						Type:    "Authentication",
						Message: "Unauthorized",
					},
				}

				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(httpError)
				return
			}

			userRepository := userrepository.New(i.Db)

			user, err := userRepository.FindByID(token.ID)

			if err != nil {
				panic(err)
			}

			i.User = domain.User{
				ID:       user.ID,
				Name:     user.Name,
				Document: user.Document,
				Email:    user.Email,
			}

			next.ServeHTTP(w, r)
		})
	}

}
