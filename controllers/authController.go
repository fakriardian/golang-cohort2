package controllers

import (
	"final-challenge/dtos"
	"final-challenge/helpers"
	"final-challenge/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authService services.AuthService
	userService services.UserService
}

func RegisterAuthController(authService services.AuthService, userService services.UserService) *authController {
	return &authController{
		authService: authService,
		userService: userService,
	}
}

func (ctrl *authController) CreateUser(ctx *gin.Context) {
	var req dtos.UserRegister

	if err := helpers.BindRequest(ctx, &req); err != nil {
		return
	}

	createUserResult, err := ctrl.userService.CreateUserService(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusInternalServerError,
			Message: "error",
		})
		return
	}

	data := []interface{}{createUserResult}
	response := dtos.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success created user",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *authController) UserLogin(ctx *gin.Context) {
	var req dtos.UserLogin

	if err := helpers.BindRequest(ctx, &req); err != nil {
		return
	}

	loginResult, err := ctrl.authService.GenerateToken(req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.Response{
			Error:   err.Error(),
			Status:  http.StatusNotFound,
			Message: "error",
		})
		return
	}

	response := dtos.Response{
		AccessToken: loginResult,
		Status:      http.StatusOK,
		Message:     "success generate token",
	}

	ctx.JSON(http.StatusOK, response)
}
