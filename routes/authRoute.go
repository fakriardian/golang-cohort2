package routes

import (
	"final-challenge/controllers"
	"final-challenge/repository"
	"final-challenge/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, router *gin.Engine) {
	userRepository := repository.RegisterUserRepository(db)
	userServices := services.RegisterUserService(userRepository)
	authServices := services.RegisterAuthService(userRepository)
	authControllers := controllers.RegisterAuthController(authServices, userServices)

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authControllers.CreateUser)
		authRouter.POST("/login", authControllers.UserLogin)
	}
}
