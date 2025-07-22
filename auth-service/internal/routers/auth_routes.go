package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRegister(router *gin.Engine, authHandler *handlers.AuthHandler) {
	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
}
