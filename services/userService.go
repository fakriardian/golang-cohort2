package services

import (
	"final-challenge/dtos"
	"final-challenge/helpers"
	"final-challenge/models"
	"final-challenge/repository"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type UserService interface {
	CreateUserService(req dtos.UserRegister) (*models.User, error)
	FindUserbyIdService(id string) (*models.User, error)
	IsExistingUserService(email string) (*models.User, error)
	UserToDto(user *models.User) *dtos.User
}

type userService struct {
	repository repository.UserRepository
}

func RegisterUserService(repository repository.UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (service *userService) CreateUserService(req dtos.UserRegister) (*models.User, error) {
	isUserExist, _ := service.IsExistingUserService(req.Email)
	if isUserExist != nil {
		return nil, fmt.Errorf(fmt.Sprintf("email %s has been Registered", req.Email))
	}

	userId := uuid.NewV4()

	input := models.User{
		ID:       userId,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	return service.repository.CreateUser(&input)
}

func (service *userService) FindUserbyIdService(ID string) (*models.User, error) {
	uuid, err := uuid.FromString(ID)
	if err != nil {
		return nil, err
	}

	return service.repository.FindbyId(uuid)
}

func (service *userService) IsExistingUserService(email string) (*models.User, error) {
	encryption := helpers.NewEncryption()
	decrypted := encryption.Encrypt(email)
	return service.repository.IsExistUser(decrypted)
}

func (service *userService) UserToDto(user *models.User) *dtos.User {
	return &dtos.User{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
