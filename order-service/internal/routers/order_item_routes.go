package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func OrderItemRegister(router *gin.Engine, ordemItemHandler *handlers.OrderItemHandler) {
	ordemItemGroup := router.Group("/order-items")
	{
		ordemItemGroup.GET("/", ordemItemHandler.GetAllOrderItems)
		ordemItemGroup.GET("/:id", ordemItemHandler.GetOrderItemByID)
	}
}
