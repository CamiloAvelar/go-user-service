package bootstrap

import (
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
	"github.com/CamiloAvelar/go-user-service/usecases/oauth"
	"github.com/CamiloAvelar/go-user-service/usecases/user"
)

func SetupRoutes(r *infrainterfaces.Router) {
	user.SetupRoutes(r)
	oauth.SetupRoutes(r)
}
