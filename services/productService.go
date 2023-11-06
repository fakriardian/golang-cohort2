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
	CreateProductService(req dtos.ProductCreateRequest) (*models.Product, error)
	UpdateProductService(productID uuid.UUID, req dtos.ProductUpdateRequest) (*models.Product, error)
	DeleteProductService(productID uuid.UUID) int
	FindProductServicebyId(id string) (*models.Product, error)
	FindProductService(filter dtos.Filter) (*[]models.Product, error)
	CountProductService(filter dtos.Filter) (*int64, error)
	PaginationProductService(filter dtos.FilterRequest) ([]interface{}, int64, int64, int64, error)
}

type service struct {
	repository repository.ProductRepository
	claudinary libs.CloudinaryService
}

func RegisterProductService(repository repository.ProductRepository, claudinary *libs.CloudinaryService) *service {
	return &service{
		repository: repository,
		claudinary: *claudinary,
	}
}

func (service *service) CreateProductService(req dtos.ProductCreateRequest) (*models.Product, error) {
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

	return service.repository.CreateProduct(&input)
}

func (service *service) FindProductServicebyId(ID string) (*models.Product, error) {
	uuid, err := uuid.FromString(ID)
	if err != nil {
		return nil, err
	}

	return service.repository.FindbyId(uuid)
}

func (service *service) UpdateProductService(productID uuid.UUID, req dtos.ProductUpdateRequest) (*models.Product, error) {
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

	return service.repository.UpdateProduct(productID, &input)
}

func (service *service) DeleteProductService(productID uuid.UUID) int {
	return service.repository.DeleteProduct(productID)
}

func (service *service) FindProductService(filter dtos.Filter) (*[]models.Product, error) {
	return service.repository.FindAll(filter)
}

func (service *service) CountProductService(filter dtos.Filter) (*int64, error) {
	return service.repository.Count(filter)
}

func (service *service) PaginationProductService(filter dtos.FilterRequest) ([]interface{}, int64, int64, int64, error) {
	filterPaging := dtos.Paginate(filter)
	var productData []models.Product

	result, err := service.FindProductService(filterPaging)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	productData = *result

	productCount, err := service.CountProductService(filterPaging)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	var results []interface{}
	for _, p := range productData {
		results = append(results, p)
	}

	return results, *productCount, int64(filterPaging.Page), int64(filterPaging.Size), nil

}
