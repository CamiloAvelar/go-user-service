package infrainterfaces

import (
	"database/sql"

	"github.com/CamiloAvelar/go-user-service/domain"
	"github.com/CamiloAvelar/go-user-service/infrastructure/config"
	"github.com/go-oauth2/oauth2/v4/server"
)

type ServerInjections struct {
	Config   config.Config
	Db       *sql.DB
	OauthSrv *server.Server
	User     domain.User
}
