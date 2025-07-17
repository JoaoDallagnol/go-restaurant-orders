package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRegister(router *gin.Engine) {
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
}
