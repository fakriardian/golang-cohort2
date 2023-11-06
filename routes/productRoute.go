package routes

import (
	"final-challenge/config"
	"final-challenge/controllers"
	"final-challenge/middlewares"
	"final-challenge/repository"
	"final-challenge/services"

	"github.com/gin-gonic/gin"
)

func InitProductRoutes(deps *config.Deps, router *gin.Engine) {
	productRepository := repository.RegisterProductRepository(deps.DB)
	productServices := services.RegisterProductService(productRepository, deps.STRG)
	productControllers := controllers.RegisterProductController(productServices)

	productRouter := router.Group("/products")
	{
		productRouter.GET("", productControllers.RetrieveProducts)
		productRouter.GET("/:id", productControllers.RetrieveProduct)

		productRouter.Use(middlewares.Authentication())
		productRouter.POST("", productControllers.CreateProduct)
		productRouter.PUT("/:id", middlewares.ProductAuthorization(deps.DB), productControllers.UpdateProduct)
		productRouter.DELETE("/:id", middlewares.ProductAuthorization(deps.DB), productControllers.DeleteProduct)

		InitVariantRoutes(deps.DB, productRouter.Group("/variants"))
	}
}
