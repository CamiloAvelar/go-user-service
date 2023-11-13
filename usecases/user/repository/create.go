package userrepository

import (
	"fmt"

	"github.com/CamiloAvelar/go-user-service/domain"
)

func (repository repository) Create(
	userRequest domain.User,
) (int64, error) {
	result, err := repository.db.Exec(
		"INSERT INTO users (name, document, email, password) VALUES (?, ?, ?, ?)",
		userRequest.Name,
		userRequest.Document,
		userRequest.Email,
		userRequest.Password,
	)

	if err != nil {
		return 0, fmt.Errorf("CreateUserRepository: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil

}
