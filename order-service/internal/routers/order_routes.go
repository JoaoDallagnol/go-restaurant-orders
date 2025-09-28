package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func OrderRegister(router *gin.Engine, orderHandler *handlers.OrderHandler) {
	orderGroup := router.Group("/orders")
	{
		orderGroup.GET("/", orderHandler.GetAllOrders)
		orderGroup.GET("/:id", orderHandler.GetOrderByID)
		orderGroup.POST("/", orderHandler.CreateOrder)
		orderGroup.PUT("/:id", orderHandler.UpdateOrder)
		orderGroup.DELETE("/:id", orderHandler.DeleteOrder)
	}
}
