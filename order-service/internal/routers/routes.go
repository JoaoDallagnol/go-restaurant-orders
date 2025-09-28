package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	orderHandler *handlers.OrderHandler,
	orderItemHandler *handlers.OrderItemHandler,
) {
	OrderRegister(router, orderHandler)
	OrderItemRegister(router, orderItemHandler)
}
