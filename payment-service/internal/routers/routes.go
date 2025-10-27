package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	paymentHandler *handlers.PaymentHandler,
) {
	PaymentRegister(router, paymentHandler)
}
