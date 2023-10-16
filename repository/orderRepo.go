package repository

import (
	"errors"
	"fmt"
	"mini-challenge/models"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(input *models.Order) (*models.Order, int)
	UpdateOrder(orderID uuid.UUID, input *models.Order) (*models.Order, int)
	DeleteOrder(orderID uuid.UUID) int
	FindAll() (*[]models.Order, error)
	FindbyId(id uuid.UUID) (*models.Order, error)
}

type repository struct {
	db *gorm.DB
}

func RegisterOrderRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateOrder(input *models.Order) (*models.Order, int) {
	tx := r.db.Begin()

	if tx.Error != nil {
		return nil, http.StatusInternalServerError
	}

	if err := tx.Create(input).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError
	}

	tx.Commit()

	return input, http.StatusCreated
}

func (r *repository) FindAll() (*[]models.Order, error) {
	var order []models.Order

	tx := r.db.Preload("Items").Find(&order)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("order data is not found")
		}
		return nil, fmt.Errorf("error finding order: %w", tx.Error)
	}

	return &order, nil
}

func (r *repository) FindbyId(ID uuid.UUID) (*models.Order, error) {
	var order models.Order

	tx := r.db.Where(models.Order{
		ID: ID,
	}).First(&order)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("order data is not found")
		}
		return nil, fmt.Errorf("error finding order: %w", tx.Error)
	}

	return &order, nil
}

func (r *repository) UpdateOrder(orderID uuid.UUID, input *models.Order) (*models.Order, int) {
	tx := r.db.Begin()

	if tx.Error != nil {
		return nil, http.StatusInternalServerError
	}

	if err := tx.Where("order_id = ?", orderID).Delete(&models.Item{}).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError
	}

	if err := tx.Model(input).Where("id = ?", orderID).Updates(input).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError
	}

	tx.Commit()

	return input, http.StatusOK
}

func (r *repository) DeleteOrder(orderID uuid.UUID) int {
	tx := r.db.Begin()

	if tx.Error != nil {
		return http.StatusInternalServerError
	}

	if err := tx.Where("order_id = ?", orderID).Delete(&models.Item{}).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError
	}

	if err := tx.Where("id = ?", orderID).Delete(&models.Order{}).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError
	}

	tx.Commit()

	return http.StatusOK
}
