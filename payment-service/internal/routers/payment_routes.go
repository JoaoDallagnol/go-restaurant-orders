package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func PaymentRegister(router *gin.Engine, paymentHandler *handlers.PaymentHandler) {
	paymentGroup := router.Group("/payments")
	{
		paymentGroup.GET("/", paymentHandler.GetAllPayments)
		paymentGroup.GET("/:id", paymentHandler.GetPaymentById)
		paymentGroup.POST("/", paymentHandler.CreatePayment)
		paymentGroup.DELETE("/:id", paymentHandler.DeletePayment)
	}
}
