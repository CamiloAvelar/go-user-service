package userusecase

import userdto "github.com/CamiloAvelar/go-user-service/usecases/user/dto"

func (usecase usecase) Create(createUserRequest userdto.CreateUserRequest) (int64, error) {
	if err := createUserRequest.Validate(); err != nil {
		return 0, err
	}

	userId, err := usecase.repository.Create(createUserRequest)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
