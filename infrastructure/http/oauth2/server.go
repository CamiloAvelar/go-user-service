package oauth2

import (
	"database/sql"

	"github.com/CamiloAvelar/go-user-service/infrastructure/config"
	"github.com/go-oauth2/mysql/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func GetServer(db *sql.DB, config config.Config) *server.Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	tokenstore := mysql.NewStoreWithDB(
		db,
		"oauth2_token",
		600,
	)

	manager.MapTokenStorage(tokenstore)
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "22222222",
		Domain: "http://localhost:9094",
		Public: true,
	})
	manager.MapClientStorage(clientStore)

	return server.NewServer(server.NewConfig(), manager)
}
