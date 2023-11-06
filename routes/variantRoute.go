package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitVariantRoutes(db *gorm.DB, variantRouter *gin.RouterGroup) {
	// userRepository := repository.RegisterUserRepository(db)
	// userServices := services.RegisterUserService(userRepository)
	// authServices := services.RegisterAuthService(userRepository)
	// authControllers := controllers.RegisterControllerAuth(authServices, userServices)

	// variantRouter := router.Group("/products")
	{
		variantRouter.GET("")
		variantRouter.GET("/:id")
		variantRouter.POST("")
		variantRouter.PUT("/:id")
		variantRouter.DELETE("/:id")
	}
}
