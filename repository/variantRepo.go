package repository

import (
	"errors"
	"final-challenge/dtos"
	"final-challenge/models"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VariantRepository interface {
	CreateVariant(input *models.Variant) (*models.Variant, error)
	UpdateVariant(variantID uuid.UUID, input *models.Variant) (*models.Variant, error)
	DeleteVariant(variantID uuid.UUID) int
	FindAll(filter dtos.Filter) ([]models.Variant, error)
	FindbyId(id uuid.UUID) (models.Variant, error)
	Count(filter dtos.Filter) (*int64, error)
}

type varinatRepository struct {
	db *gorm.DB
}

func RegisterVariantRepository(db *gorm.DB) *varinatRepository {
	return &varinatRepository{
		db: db,
	}
}

func (r *varinatRepository) CreateVariant(input *models.Variant) (*models.Variant, error) {
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

func (r *varinatRepository) FindAll(filter dtos.Filter) ([]models.Variant, error) {
	var variant []models.Variant

	tx := r.db

	if filter.Search != "" {
		tx = tx.Where("name LIKE ?", filter.Search)
	}

	tx = tx.Preload(clause.Associations).Limit(filter.Size).Offset((filter.Page - 1) * filter.Size).Find(&variant)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("variant data is not found")
		}
		return nil, fmt.Errorf("error finding variant: %w", tx.Error)
	}

	return variant, nil
}

func (r *varinatRepository) FindbyId(ID uuid.UUID) (models.Variant, error) {
	var variant models.Variant

	tx := r.db.Preload(clause.Associations).Where(models.Variant{
		ID: ID,
	}).First(&variant)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return variant, errors.New("variant data is not found")
		}
		return variant, fmt.Errorf("error finding variant: %w", tx.Error)
	}

	return variant, nil
}

func (r *varinatRepository) UpdateVariant(variantID uuid.UUID, input *models.Variant) (*models.Variant, error) {
	tx := r.db.Begin()

	if tx.Error != nil {
		return nil, errors.New("error when Updated Data")
	}

	if err := tx.Model(input).Where("id = ?", variantID).Updates(input).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error when Updated Data")
	}

	tx.Commit()

	var variant models.Variant
	if err := r.db.First(&variant, variantID).Error; err != nil {
		return nil, errors.New("error when Updated Data")
	}

	return &variant, nil
}

func (r *varinatRepository) DeleteVariant(variantID uuid.UUID) int {
	tx := r.db.Begin()

	if tx.Error != nil {
		return http.StatusInternalServerError
	}

	if err := tx.Where("id = ?", variantID).Delete(&models.Variant{}).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError
	}

	tx.Commit()

	return http.StatusOK
}

func (r *varinatRepository) Count(filter dtos.Filter) (*int64, error) {
	var total int64

	tx := r.db.Model(&models.Variant{})
	if filter.Search != "" {
		tx = tx.Where("name LIKE ?", filter.Search)
	}
	tx = tx.Count(&total)

	if tx.Error != nil {
		return nil, fmt.Errorf("error count variant: %w", tx.Error)
	}

	return &total, nil
}
