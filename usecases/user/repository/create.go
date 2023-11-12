package userrepository

import (
	"fmt"

	userdto "github.com/CamiloAvelar/go-user-service/usecases/user/dto"
)

func (repository repository) Create(
	userRequest userdto.CreateUserRequest,
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
