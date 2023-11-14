package infrainterfaces

import (
	"database/sql"

	"github.com/CamiloAvelar/go-user-service/domain"
	"github.com/CamiloAvelar/go-user-service/infrastructure/config"
)

type ServerInjections struct {
	Config config.Config
	Db     *sql.DB
	User   domain.User
}
