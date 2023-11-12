package userusecase

import (
	"github.com/CamiloAvelar/go-user-service/domain/errors"
	userdto "github.com/CamiloAvelar/go-user-service/usecases/user/dto"
)

func (usecase usecase) Login(loginUserRequest userdto.LoginUserRequest) (userdto.LoginUserResponse, error) {
	if err := loginUserRequest.Validate(); err != nil {
		return userdto.LoginUserResponse{}, err
	}

	user, err := usecase.repository.
		FindByEmailOrDocument(
			loginUserRequest.Email,
			loginUserRequest.Document,
		)

	if err != nil {
		return userdto.LoginUserResponse{}, err
	}

	if user.ID == 0 {
		return userdto.LoginUserResponse{}, &errors.ApplicationError{
			Type:       "Invalid",
			Message:    "Invalid credentials",
			StatusCode: 400,
		}
	}

	if ok := loginUserRequest.ComparePassword(user.Password); !ok {
		return userdto.LoginUserResponse{}, &errors.ApplicationError{
			Type:       "Invalid",
			Message:    "Invalid credentials",
			StatusCode: 400,
		}
	}

	token, err := loginUserRequest.
		CreateAccessToken(usecase.config.AccessTokenSecret, usecase.config.AccessTokenExp)

	if err != nil {
		return userdto.LoginUserResponse{}, err
	}

	return userdto.LoginUserResponse{
		AccessToken: token,
	}, nil
}
