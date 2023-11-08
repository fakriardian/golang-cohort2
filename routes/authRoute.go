package routes

import (
	"final-challenge/config"
	"final-challenge/controllers"
	"final-challenge/repository"
	"final-challenge/services"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(deps *config.Deps, router *gin.Engine) {
	userRepository := repository.RegisterUserRepository(deps.DB)
	userServices := services.RegisterUserService(userRepository)
	authServices := services.RegisterAuthService(userServices, deps.STRG)
	authControllers := controllers.RegisterAuthController(authServices, userServices)

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authControllers.CreateUser)
		authRouter.POST("/login", authControllers.UserLogin)
	}
}
