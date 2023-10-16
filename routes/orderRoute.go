package routes

import (
	"mini-challenge/controllers"
	"mini-challenge/repository"
	"mini-challenge/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitOrderRoutes(db *gorm.DB, router *gin.Engine) {
	orderRepository := repository.RegisterOrderRepository(db)
	orderServices := services.RegisterOrderService(orderRepository)
	orderControllers := controllers.RegisterControllerOrder(orderServices)

	route := router.Group("orders")
	route.POST("/", orderControllers.CreateOrder)
	route.PUT("/:id", orderControllers.UpdateOrder)
	route.GET("/", orderControllers.RetrieveOrders)
	route.DELETE("/:id", orderControllers.DeleteOrder)
}
