package controllers

import (
	"mini-challenge/dtos"
	"mini-challenge/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service services.OrderService
}

func RegisterControllerOrder(service services.OrderService) *controller {
	return &controller{
		service: service,
	}
}

func (ctrl *controller) CreateOrder(ctx *gin.Context) {
	var req dtos.OrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

	createOrderResult, statusCode := ctrl.service.CreateOrderService(req)

	if statusCode != http.StatusCreated {
		ctx.JSON(http.StatusInternalServerError, dtos.Response{
			Error:   "Error when Inserted Data",
			Status:  http.StatusInternalServerError,
			Message: "error",
		})
		return
	}

	data := []interface{}{createOrderResult}
	response := dtos.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success created order data",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *controller) UpdateOrder(ctx *gin.Context) {
	var req dtos.OrderRequest
	orderId := ctx.Param("id")

	if orderId == "" {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   "Order ID not found",
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

	existingOrder, err := ctrl.service.FindOrderServicebyId(orderId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	updateOrderResult, statusCode := ctrl.service.UpdateOrderService(existingOrder.ID, req)

	if statusCode != http.StatusOK {
		ctx.JSON(http.StatusInternalServerError, dtos.Response{
			Error:   "Error when Updated Data",
			Status:  http.StatusInternalServerError,
			Message: "error",
		})
		return
	}

	data := []interface{}{updateOrderResult}
	response := dtos.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success updated order data",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *controller) DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("id")

	if orderId == "" {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   "Order ID is not found",
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

	existingOrder, err := ctrl.service.FindOrderServicebyId(orderId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	statusCode := ctrl.service.DeleteOrderService(existingOrder.ID)

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

func (ctrl *controller) RetrieveOrders(ctx *gin.Context) {

	retrieveOrder, err := ctrl.service.FindOrderService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusInternalServerError,
			Message: "error",
		})
		return
	}

	data := []interface{}{retrieveOrder}
	response := dtos.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success get order datas",
	}

	ctx.JSON(http.StatusOK, response)
}
