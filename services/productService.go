package services

import (
	"final-challenge/dtos"
	"final-challenge/libs"
	"final-challenge/models"
	"final-challenge/repository"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type ProductService interface {
	CreateProductService(req dtos.ProductCreateRequest) (*dtos.Product, error)
	UpdateProductService(productID uuid.UUID, req dtos.ProductUpdateRequest) (*dtos.Product, error)
	DeleteProductService(productID uuid.UUID) int
	FindProductServicebyId(id string) (models.Product, error)
	FindProductService(filter dtos.Filter) ([]models.Product, error)
	CountProductService(filter dtos.Filter) (*int64, error)
	PaginationProductService(filter dtos.FilterRequest) ([]interface{}, int64, int64, int64, error)
	RetrieveProductService(id string) ([]interface{}, error)
}

type service struct {
	repository     repository.ProductRepository
	variantService VariantService
	userService    UserService
	claudinary     libs.CloudinaryService
}

func RegisterProductService(repository repository.ProductRepository, variantService VariantService, userService UserService, claudinary *libs.CloudinaryService) *service {
	return &service{
		repository:     repository,
		variantService: variantService,
		userService:    userService,
		claudinary:     *claudinary,
	}
}

func (service *service) CreateProductService(req dtos.ProductCreateRequest) (*dtos.Product, error) {
	productId := uuid.NewV4()
	userID, _ := uuid.FromString(req.AdminID)

	resultUploadFile, err := service.claudinary.UploadFile(req.File, req.File.Filename)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("can't upload file: %s", err))
	}

	input := models.Product{
		ID:       productId,
		Name:     req.Name,
		ImageUrl: resultUploadFile,
		UserID:   userID,
	}

	result, err := service.repository.CreateProduct(&input)
	if err != nil {
		return nil, err
	}

	return productToDto(result), nil
}

func (service *service) FindProductServicebyId(ID string) (models.Product, error) {
	uuid, err := uuid.FromString(ID)
	if err != nil {
		return models.Product{}, err
	}

	return service.repository.FindbyId(uuid)
}

func (service *service) UpdateProductService(productID uuid.UUID, req dtos.ProductUpdateRequest) (*dtos.Product, error) {
	userID, _ := uuid.FromString(req.AdminID)

	input := models.Product{
		ID:     productID,
		Name:   req.Name,
		UserID: userID,
	}

	if req.File != nil {
		resultUploadFile, err := service.claudinary.UploadFile(req.File, req.File.Filename)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("can't upload file: %s", err))
		}
		input.ImageUrl = resultUploadFile
	}

	result, err := service.repository.UpdateProduct(productID, &input)
	if err != nil {
		return nil, err
	}

	return productToDto(result), nil

}

func (service *service) DeleteProductService(productID uuid.UUID) int {
	return service.repository.DeleteProduct(productID)
}

func (service *service) FindProductService(filter dtos.Filter) ([]models.Product, error) {
	return service.repository.FindAll(filter)
}

func (service *service) CountProductService(filter dtos.Filter) (*int64, error) {
	return service.repository.Count(filter)
}

func (service *service) RetrieveProductService(id string) ([]interface{}, error) {
	result, err := service.FindProductServicebyId(id)
	if err != nil {
		return nil, err
	}

	results := []interface{}{service.productListDto(result)}

	return results, nil
}

func (service *service) PaginationProductService(filter dtos.FilterRequest) ([]interface{}, int64, int64, int64, error) {
	filterPaging := dtos.Paginate(filter)

	result, err := service.FindProductService(filterPaging)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	productCount, err := service.CountProductService(filterPaging)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	results := make([]interface{}, len(result))
	for i, v := range result {
		results[i] = service.productListDto(v)
	}

	return results, *productCount, int64(filterPaging.Page), int64(filterPaging.Size), nil

}

func productToDto(result *models.Product) *dtos.Product {
	response := &dtos.Product{
		ID:        result.ID.String(),
		Name:      result.Name,
		ImageUrl:  result.ImageUrl,
		UserID:    result.UserID.String(),
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	return response
}

func (service *service) productListDto(result models.Product) dtos.ProductListResponse {
	var variantData []dtos.Variant
	for _, v := range result.Variants {
		variant := *service.variantService.VariantToDto(&v)
		variantData = append(variantData, variant)
	}

	userData := service.userService.UserToDto(&result.User)

	return dtos.ProductListResponse{
		ID:        result.ID.String(),
		Name:      result.Name,
		ImageUrl:  result.ImageUrl,
		UserID:    *userData,
		Variants:  variantData,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

}
