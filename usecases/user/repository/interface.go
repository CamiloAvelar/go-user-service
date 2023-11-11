package userrepository

import (
	"database/sql"

	"github.com/CamiloAvelar/go-user-service/domain"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) domain.UserRepository {
	return &repository{
		db: db,
	}
}
