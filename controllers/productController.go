package controllers

import (
	"final-challenge/dtos"
	"final-challenge/helpers"
	"final-challenge/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service services.ProductService
}

func RegisterProductController(service services.ProductService) *controller {
	return &controller{
		service: service,
	}
}

func (ctrl *controller) CreateProduct(ctx *gin.Context) {
	var req dtos.ProductCreateRequest

	adminID := ctx.Request.Header.Get("AdminID")
	req.AdminID = adminID

	if err := helpers.BindRequest(ctx, &req); err != nil {
		return
	}

	createProductResult, err := ctrl.service.CreateProductService(req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

	data := []interface{}{createProductResult}
	response := dtos.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success created product data",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *controller) UpdateProduct(ctx *gin.Context) {
	var req dtos.ProductUpdateRequest
	var uri dtos.ProductIDUri

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	adminID := ctx.Request.Header.Get("AdminID")
	req.AdminID = adminID

	if err := helpers.BindRequest(ctx, &req); err != nil {
		return
	}

	existingProduct, err := ctrl.service.FindProductServicebyId(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	updateProductResult, err := ctrl.service.UpdateProductService(existingProduct.ID, req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

	data := []interface{}{updateProductResult}
	response := dtos.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success updated product data",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *controller) DeleteProduct(ctx *gin.Context) {
	var uri dtos.ProductIDUri

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	existingProduct, err := ctrl.service.FindProductServicebyId(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	statusCode := ctrl.service.DeleteProductService(existingProduct.ID)

	if statusCode != http.StatusOK {
		ctx.JSON(http.StatusInternalServerError, dtos.Response{
			Error:   "Error when Deleted Data",
			Status:  http.StatusInternalServerError,
			Message: "error",
		})
		return
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "success deleted data",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *controller) RetrieveProducts(ctx *gin.Context) {
	var filter dtos.FilterRequest

	if err := helpers.BindRequest(ctx, &filter); err != nil {
		return
	}

	retrieveProduct, total, page, size, err := ctrl.service.PaginationProductService(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusInternalServerError,
			Message: "error",
		})
		return
	}

	if retrieveProduct == nil {
		retrieveProduct = []interface{}{}
	}

	response := dtos.PaginationResponse{
		Data:     retrieveProduct,
		Total:    total,
		Page:     page,
		PageSize: size,
		Status:   http.StatusOK,
		Message:  "success retrieve products data",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *controller) RetrieveProduct(ctx *gin.Context) {
	var uri dtos.ProductIDUri

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	data, err := ctrl.service.RetrieveProductService(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	response := dtos.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success retrieve product data",
	}

	ctx.JSON(http.StatusOK, response)
}
