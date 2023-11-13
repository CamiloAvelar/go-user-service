package userusecase

import (
	"github.com/CamiloAvelar/go-user-service/domain"
	"github.com/CamiloAvelar/go-user-service/domain/errors"
)

func (usecase usecase) Login(loginUserRequest domain.Login) (*domain.LoginResponse, error) {
	if err := loginUserRequest.Validate(); err != nil {
		return nil, err
	}

	user, err := usecase.repository.
		FindByEmailOrDocument(
			loginUserRequest.Email,
			loginUserRequest.Document,
		)

	if err != nil {
		return nil, err
	}

	if err := user.ValidatePersisted(); err != nil {
		return nil, &errors.ApplicationError{
			Type:       "Invalid",
			Message:    "Invalid credentials",
			StatusCode: 400,
		}
	}

	if ok := loginUserRequest.ComparePassword(user.Password); !ok {
		return nil, &errors.ApplicationError{
			Type:       "Invalid",
			Message:    "Invalid credentials",
			StatusCode: 400,
		}
	}

	loginUser := domain.Login{User: user}

	token, err := loginUser.
		CreateAccessToken(usecase.config.AccessTokenSecret, usecase.config.AccessTokenExp)

	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		AccessToken: token,
	}, nil
}
