package services

import (
	"errors"
	"final-challenge/dtos"
	"final-challenge/helpers"
	"final-challenge/libs"
	"fmt"
)

type AuthService interface {
	GenerateToken(req dtos.UserLogin) (string, error)
}

type authservice struct {
	userService UserService
	claudinary  libs.CloudinaryService
}

func RegisterAuthService(userService UserService, claudinary *libs.CloudinaryService) *authservice {
	return &authservice{
		userService: userService,
		claudinary:  *claudinary,
	}
}

func (service *authservice) GenerateToken(req dtos.UserLogin) (string, error) {
	isUserExist, _ := service.userService.IsExistingUserService(req.Email)
	if isUserExist == nil {
		return "", fmt.Errorf("email %s hasn't been Registered", req.Email)
	}

	validPass, err := helpers.ComparePassword(req.Password, isUserExist.Password)
	if err != nil || !validPass {
		return "", errors.New("unauthorized")
	}

	getPrivateKey, err := service.claudinary.Getkey("private")
	if err != nil {
		return "", errors.New("unauthorized")
	}

	return helpers.GenerateAccessToken(isUserExist.ID.String(), getPrivateKey)
}
