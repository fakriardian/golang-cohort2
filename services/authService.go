package services

import (
	"errors"
	"final-challenge/dtos"
	"final-challenge/helpers"
	"final-challenge/repository"
	"fmt"
)

type AuthService interface {
	GenerateToken(req dtos.UserLogin) (string, error)
}

func RegisterAuthService(repository repository.UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (service *userService) GenerateToken(req dtos.UserLogin) (string, error) {
	isUserExist, _ := service.IsExistingUserService(req.Email)
	if isUserExist == nil {
		return "", fmt.Errorf("email %s hasn't been Registered", req.Email)
	}

	validPass, err := helpers.ComparePassword(req.Password, isUserExist.Password)
	if err != nil || !validPass {
		return "", errors.New("unauthorized")
	}

	return helpers.GenerateAccessToken(isUserExist.ID.String())
}
