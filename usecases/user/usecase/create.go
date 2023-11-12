package userusecase

import (
	"github.com/CamiloAvelar/go-user-service/domain/errors"
	userdto "github.com/CamiloAvelar/go-user-service/usecases/user/dto"
)

func (usecase usecase) Create(createUserRequest userdto.CreateUserRequest) (int64, error) {
	if err := createUserRequest.Validate(); err != nil {
		return 0, err
	}

	user, err := usecase.repository.
		FindByEmailOrDocument(
			createUserRequest.Email,
			createUserRequest.Document,
		)

	if err != nil {
		return 0, err
	}

	if user.ID != 0 {
		return 0, &errors.ApplicationError{
			Type:       "Duplicated",
			Message:    "User already exists",
			StatusCode: 400,
		}
	}

	if err := createUserRequest.EncryptPassword(); err != nil {
		return 0, err
	}

	userId, err := usecase.repository.Create(createUserRequest)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
