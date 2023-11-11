package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CamiloAvelar/go-user-service/infrastructure/config"
	"github.com/go-sql-driver/mysql"
)

func GetConnection(c config.Config) *sql.DB {
	cfg := mysql.Config{
		User:   c.DBUser,
		Passwd: c.DBPass,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", c.DBHost, c.DBPort),
		DBName: c.DBName,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
