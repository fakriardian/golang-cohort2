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

type ProductRepository interface {
	CreateProduct(input *models.Product) (*models.Product, error)
	UpdateProduct(productID uuid.UUID, input *models.Product) (*models.Product, error)
	DeleteProduct(productID uuid.UUID) int
	FindAll(filter dtos.Filter) (*[]models.Product, error)
	FindbyId(id uuid.UUID) (*models.Product, error)
	Count(filter dtos.Filter) (*int64, error)
}

type repository struct {
	db *gorm.DB
}

func RegisterProductRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateProduct(input *models.Product) (*models.Product, error) {
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

func (r *repository) FindAll(filter dtos.Filter) (*[]models.Product, error) {
	var product []models.Product

	tx := r.db

	if filter.Search != "" {
		tx = tx.Where("name LIKE ?", filter.Search)
	}

	tx = tx.Preload(clause.Associations).Limit(filter.Size).Offset((filter.Page - 1) * filter.Size).Find(&product)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product data is not found")
		}
		return nil, fmt.Errorf("error finding product: %w", tx.Error)
	}

	return &product, nil
}

func (r *repository) FindbyId(ID uuid.UUID) (*models.Product, error) {
	var product models.Product

	tx := r.db.Preload(clause.Associations).Where(models.Product{
		ID: ID,
	}).First(&product)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product data is not found")
		}
		return nil, fmt.Errorf("error finding product: %w", tx.Error)
	}

	return &product, nil
}

func (r *repository) UpdateProduct(productID uuid.UUID, input *models.Product) (*models.Product, error) {
	tx := r.db.Begin()

	if tx.Error != nil {
		return nil, errors.New("error when Updated Data")
	}

	if err := tx.Where("product_id = ?", productID).Delete(&models.Variant{}).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error when Inserted Data")
	}

	if err := tx.Model(input).Where("id = ?", productID).Updates(input).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error when Inserted Data")
	}

	tx.Commit()

	return input, nil
}

func (r *repository) DeleteProduct(productID uuid.UUID) int {
	tx := r.db.Begin()

	if tx.Error != nil {
		return http.StatusInternalServerError
	}

	if err := tx.Where("product_id = ?", productID).Delete(&models.Variant{}).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError
	}

	if err := tx.Where("id = ?", productID).Delete(&models.Product{}).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError
	}

	tx.Commit()

	return http.StatusOK
}

func (r *repository) Count(filter dtos.Filter) (*int64, error) {
	var total int64

	tx := r.db.Model(&models.Product{})
	if filter.Search != "" {
		tx = tx.Where("name LIKE ?", filter.Search)
	}
	tx = tx.Count(&total)

	if tx.Error != nil {
		return nil, fmt.Errorf("error count product: %w", tx.Error)
	}

	return &total, nil
}
