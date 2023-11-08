package services

import (
	"final-challenge/dtos"
	"final-challenge/models"
	"final-challenge/repository"

	uuid "github.com/satori/go.uuid"
)

type VariantService interface {
	CreateVariantService(req dtos.VariantCreateRequest) (*models.Variant, error)
	UpdateVariantService(variantID uuid.UUID, req dtos.VariantUpdateRequest) (*models.Variant, error)
	DeleteVariantService(variantID uuid.UUID) int
	FindVariantServicebyId(id string) (*models.Variant, error)
	FindVariantService(filter dtos.Filter) (*[]models.Variant, error)
	CountVariantService(filter dtos.Filter) (*int64, error)
	PaginationVariantService(filter dtos.FilterRequest) ([]interface{}, int64, int64, int64, error)
}

type variantService struct {
	repository repository.VariantRepository
}

func RegisterVariantService(repository repository.VariantRepository) *variantService {
	return &variantService{
		repository: repository,
	}
}

func (service *variantService) CreateVariantService(req dtos.VariantCreateRequest) (*models.Variant, error) {
	variantId := uuid.NewV4()
	productID, _ := uuid.FromString(req.ProductID)

	input := models.Variant{
		ID:        variantId,
		Name:      req.Name,
		Quantity:  req.Quantity,
		ProductID: productID,
	}

	return service.repository.CreateVariant(&input)
}

func (service *variantService) FindVariantServicebyId(ID string) (*models.Variant, error) {
	uuid, err := uuid.FromString(ID)
	if err != nil {
		return nil, err
	}

	return service.repository.FindbyId(uuid)
}

func (service *variantService) UpdateVariantService(variantID uuid.UUID, req dtos.VariantUpdateRequest) (*models.Variant, error) {
	// userID, _ := uuid.FromString(req.AdminID)

	input := models.Variant{
		ID:       variantID,
		Name:     req.Name,
		Quantity: req.Quantity,
		// UserID: userID,
	}

	return service.repository.UpdateVariant(variantID, &input)
}

func (service *variantService) DeleteVariantService(variantID uuid.UUID) int {
	return service.repository.DeleteVariant(variantID)
}

func (service *variantService) FindVariantService(filter dtos.Filter) (*[]models.Variant, error) {
	return service.repository.FindAll(filter)
}

func (service *variantService) CountVariantService(filter dtos.Filter) (*int64, error) {
	return service.repository.Count(filter)
}

func (service *variantService) PaginationVariantService(filter dtos.FilterRequest) ([]interface{}, int64, int64, int64, error) {
	filterPaging := dtos.Paginate(filter)
	var variantData []models.Variant

	result, err := service.FindVariantService(filterPaging)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	variantData = *result

	variantCount, err := service.CountVariantService(filterPaging)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	var results []interface{}
	for _, p := range variantData {
		results = append(results, p)
	}

	return results, *variantCount, int64(filterPaging.Page), int64(filterPaging.Size), nil

}
