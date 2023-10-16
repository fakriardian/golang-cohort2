package services

import (
	"mini-challenge/dtos"
	"mini-challenge/models"
	"mini-challenge/repository"
	"time"

	uuid "github.com/satori/go.uuid"
)

type OrderService interface {
	CreateOrderService(req dtos.OrderRequest) (*models.Order, int)
	UpdateOrderService(orderID uuid.UUID, req dtos.OrderRequest) (*models.Order, int)
	DeleteOrderService(orderID uuid.UUID) int
	FindOrderService() (*[]models.Order, error)
	FindOrderServicebyId(id string) (*models.Order, error)
}

type service struct {
	repository repository.OrderRepository
}

func RegisterOrderService(repository repository.OrderRepository) *service {
	return &service{
		repository: repository,
	}
}

func (service *service) CreateOrderService(req dtos.OrderRequest) (*models.Order, int) {
	orderId := uuid.NewV4()
	var items []models.Item
	orderedAt, _ := time.Parse(time.RFC3339, req.OrderedAt)

	for _, value := range req.Items {
		item := models.Item{
			ID:          uuid.NewV4(),
			Name:        value.Name,
			Description: value.Description,
			Quantity:    value.Quantity,
			OrderID:     orderId,
		}
		items = append(items, item)
	}

	input := models.Order{
		ID:           orderId,
		CustomerName: req.CustomerName,
		OrederedAt:   orderedAt,
		Items:        items,
	}

	return service.repository.CreateOrder(&input)
}

func (service *service) FindOrderService() (*[]models.Order, error) {
	return service.repository.FindAll()
}

func (service *service) FindOrderServicebyId(ID string) (*models.Order, error) {
	uuid, err := uuid.FromString(ID)
	if err != nil {
		return nil, err
	}

	return service.repository.FindbyId(uuid)
}

func (service *service) UpdateOrderService(orderID uuid.UUID, req dtos.OrderRequest) (*models.Order, int) {
	var items []models.Item
	orderedAt, _ := time.Parse(time.RFC3339, req.OrderedAt)

	for _, value := range req.Items {
		item := models.Item{
			ID:          uuid.NewV4(),
			Name:        value.Name,
			Description: value.Description,
			Quantity:    value.Quantity,
			OrderID:     orderID,
		}
		items = append(items, item)
	}

	input := models.Order{
		ID:           orderID,
		CustomerName: req.CustomerName,
		OrederedAt:   orderedAt,
		Items:        items,
	}

	return service.repository.UpdateOrder(orderID, &input)
}

func (service *service) DeleteOrderService(orderID uuid.UUID) int {
	return service.repository.DeleteOrder(orderID)
}
