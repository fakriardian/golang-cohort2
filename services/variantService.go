package services

import (
	"final-challenge/dtos"
	"final-challenge/models"
	"final-challenge/repository"
	"time"

	uuid "github.com/satori/go.uuid"
)

type VariantService interface {
	CreateVariantService(req dtos.VariantCreateRequest) (*dtos.Variant, error)
	UpdateVariantService(variantID uuid.UUID, req dtos.VariantUpdateRequest) (*dtos.Variant, error)
	DeleteVariantService(variantID uuid.UUID) int
	FindVariantServicebyId(id string) (models.Variant, error)
	FindVariantService(filter dtos.Filter) ([]models.Variant, error)
	CountVariantService(filter dtos.Filter) (*int64, error)
	PaginationVariantService(filter dtos.FilterRequest) ([]interface{}, int64, int64, int64, error)
	RetrieveVariantService(id string) ([]interface{}, error)
	variantListDto(result models.Variant) dtos.VariantListResponse
	VariantToDto(result *models.Variant) *dtos.Variant
}

type variantService struct {
	repository        repository.VariantRepository
	productRepository repository.ProductRepository
}

func RegisterVariantService(repository repository.VariantRepository, productRepository repository.ProductRepository) *variantService {
	return &variantService{
		repository:        repository,
		productRepository: productRepository,
	}
}

func (service *variantService) CreateVariantService(req dtos.VariantCreateRequest) (*dtos.Variant, error) {
	variantId := uuid.NewV4()
	productID, _ := uuid.FromString(req.ProductID)

	input := models.Variant{
		ID:        variantId,
		Name:      req.Name,
		Quantity:  req.Quantity,
		ProductID: productID,
	}

	result, err := service.repository.CreateVariant(&input)
	if err != nil {
		return nil, err
	}

	return service.VariantToDto(result), nil

}

func (service *variantService) UpdateVariantService(variantID uuid.UUID, req dtos.VariantUpdateRequest) (*dtos.Variant, error) {
	input := models.Variant{
		ID:       variantID,
		Name:     req.Name,
		Quantity: req.Quantity,
	}

	result, err := service.repository.UpdateVariant(variantID, &input)
	if err != nil {
		return nil, err
	}

	return service.VariantToDto(result), nil
}

func (service *variantService) FindVariantServicebyId(ID string) (models.Variant, error) {
	uuid, err := uuid.FromString(ID)
	if err != nil {
		return models.Variant{}, err
	}

	return service.repository.FindbyId(uuid)
}

func (service *variantService) DeleteVariantService(variantID uuid.UUID) int {
	return service.repository.DeleteVariant(variantID)
}

func (service *variantService) FindVariantService(filter dtos.Filter) ([]models.Variant, error) {
	return service.repository.FindAll(filter)
}

func (service *variantService) CountVariantService(filter dtos.Filter) (*int64, error) {
	return service.repository.Count(filter)
}

func (service *variantService) PaginationVariantService(filter dtos.FilterRequest) ([]interface{}, int64, int64, int64, error) {
	filterPaging := dtos.Paginate(filter)

	result, err := service.FindVariantService(filterPaging)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	variantCount, err := service.CountVariantService(filterPaging)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	results := make([]interface{}, len(result))
	for i, v := range result {
		results[i] = service.variantListDto(v)
	}

	return results, *variantCount, int64(filterPaging.Page), int64(filterPaging.Size), nil
}

func (service *variantService) RetrieveVariantService(id string) ([]interface{}, error) {
	result, err := service.FindVariantServicebyId(id)
	if err != nil {
		return nil, err
	}

	results := []interface{}{service.variantListDto(result)}

	return results, nil
}

func (service *variantService) VariantToDto(result *models.Variant) *dtos.Variant {
	response := &dtos.Variant{
		ID:        result.ID.String(),
		Name:      result.Name,
		Quantity:  result.Quantity,
		ProductID: result.ProductID.String(),
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	return response
}

func (service *variantService) variantListDto(result models.Variant) dtos.VariantListResponse {
	productData, _ := service.productRepository.FindbyId(result.ProductID)
	return dtos.VariantListResponse{
		ID:       result.ID.String(),
		Name:     result.Name,
		Quantity: result.Quantity,
		Product: struct {
			ID        string     `json:"id,omitempty"`
			Name      string     `json:"name"`
			ImageUrl  string     `json:"imageUrl"`
			UserID    string     `json:"adminId"`
			CreatedAt *time.Time `json:"createdAt"`
			UpdatedAt *time.Time `json:"updatedAt"`
		}{
			ID:        productData.ID.String(),
			Name:      productData.Name,
			ImageUrl:  productData.ImageUrl,
			UserID:    productData.UserID.String(),
			CreatedAt: productData.CreatedAt,
			UpdatedAt: productData.UpdatedAt,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
}
