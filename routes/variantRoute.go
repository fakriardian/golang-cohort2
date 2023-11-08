package routes

import (
	"final-challenge/config"
	"final-challenge/controllers"
	"final-challenge/middlewares"
	"final-challenge/repository"
	"final-challenge/services"

	"github.com/gin-gonic/gin"
)

func InitVariantRoutes(deps *config.Deps, variantRouter *gin.RouterGroup) {
	variantRepository := repository.RegisterVariantRepository(deps.DB)
	productRepository := repository.RegisterProductRepository(deps.DB)
	variantServices := services.RegisterVariantService(variantRepository, productRepository)
	variantControllers := controllers.RegisterVariantController(variantServices)

	{
		variantRouter.GET("", variantControllers.RetrieveVariants)
		variantRouter.GET("/:id", variantControllers.RetrieveVariant)

		variantRouter.Use(middlewares.Authentication(deps.STRG))
		variantRouter.POST("", middlewares.VariantCreateAuthorization(deps.DB), variantControllers.CreateVariant)
		variantRouter.PUT("/:id", middlewares.VariantAuthorization(deps.DB), variantControllers.UpdateVariant)
		variantRouter.DELETE("/:id", middlewares.VariantAuthorization(deps.DB), variantControllers.DeleteVariant)
	}
}
