package repository

import (
	"errors"
	"final-challenge/models"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(input *models.User) (*models.User, error)
	FindbyId(id uuid.UUID) (*models.User, error)
	IsExistUser(Email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func RegisterUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(input *models.User) (*models.User, error) {
	tx := r.db.Begin()

	if tx.Error != nil {
		return nil, errors.New("error when Inserted Data")
	}

	if err := tx.Create(input).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error when Inserted Data")
	}

	tx.Commit()

	return input, nil
}

func (r *userRepository) FindbyId(ID uuid.UUID) (*models.User, error) {
	var user models.User

	tx := r.db.Where(models.User{
		ID: ID,
	}).First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user data is not found")
		}
		return nil, fmt.Errorf("error finding user: %w", tx.Error)
	}

	return &user, nil
}

func (r *userRepository) IsExistUser(Email string) (*models.User, error) {
	var user models.User

	tx := r.db.Where(models.User{
		Email: Email,
	}).First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user data is not found")
		}
		return nil, fmt.Errorf("error finding user: %w", tx.Error)
	}

	return &user, nil
}
