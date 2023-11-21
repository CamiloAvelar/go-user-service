package session

import (
	"database/sql"

	sessionmysql "github.com/go-session/mysql"
	"github.com/go-session/session"
)

func SetStore(db *sql.DB) {
	session.InitManager(
		session.SetStore(sessionmysql.NewStoreWithDB(db, "session", 600)),
	)
}
