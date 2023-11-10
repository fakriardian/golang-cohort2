package controllers

import (
	"final-challenge/dtos"
	"final-challenge/helpers"
	"final-challenge/services"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type variantController struct {
	service services.VariantService
}

func RegisterVariantController(service services.VariantService) *variantController {
	return &variantController{
		service: service,
	}
}

func (ctrl *variantController) CreateVariant(ctx *gin.Context) {
	var req dtos.VariantCreateRequest

	adminID := ctx.Request.Header.Get("AdminID")
	req.AdminID = adminID

	if err := helpers.BindRequest(ctx, &req); err != nil {
		return
	}

	createVariantResult, err := ctrl.service.CreateVariantService(req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

	data := []interface{}{createVariantResult}
	response := dtos.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success created variant data",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *variantController) UpdateVariant(ctx *gin.Context) {
	var req dtos.VariantUpdateRequest
	var uri dtos.VariantIDUri

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

	existingVariant, err := ctrl.service.FindVariantServicebyId(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	updateVariantResult, err := ctrl.service.UpdateVariantService(existingVariant.ID, req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

	data := []interface{}{updateVariantResult}
	response := dtos.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success updated variant data",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *variantController) DeleteVariant(ctx *gin.Context) {
	var uri dtos.VariantIDUri

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	existingVariant, err := ctrl.service.FindVariantServicebyId(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	statusCode := ctrl.service.DeleteVariantService(existingVariant.ID)

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

func (ctrl *variantController) RetrieveVariants(ctx *gin.Context) {
	var filter dtos.FilterRequest

	if err := helpers.BindRequest(ctx, &filter); err != nil {
		return
	}

	retrieveVariant, total, page, size, err := ctrl.service.PaginationVariantService(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusInternalServerError,
			Message: "error",
		})
		return
	}

	if retrieveVariant == nil {
		retrieveVariant = []interface{}{}
	}

	response := dtos.PaginationResponse{
		Data:      retrieveVariant,
		Total:     total,
		Page:      page,
		Size:      size,
		TotalPage: int64(math.Ceil(float64(total) / float64(size))),
		Status:    http.StatusOK,
		Message:   "success retrieve variants data",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *variantController) RetrieveVariant(ctx *gin.Context) {
	var uri dtos.VariantIDUri

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	data, err := ctrl.service.RetrieveVariantService(uri.ID)
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
		Message: "success retrieve variant data",
	}

	ctx.JSON(http.StatusOK, response)
}
