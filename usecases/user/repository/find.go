package userrepository

import (
	"database/sql"
	"errors"

	"github.com/CamiloAvelar/go-user-service/domain"
)

func (repository repository) FindByEmailOrDocument(
	email string,
	document string,
) (domain.User, error) {
	var user domain.User

	err := repository.db.
		QueryRow("SELECT id, name, document, email, password FROM users WHERE email = ? OR document = ?",
			email,
			document,
		).
		Scan(&user.ID, &user.Name, &user.Document, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, nil
		}

		return domain.User{}, err
	}

	return user, nil
}
