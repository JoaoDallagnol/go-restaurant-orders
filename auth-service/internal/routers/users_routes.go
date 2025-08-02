package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/handlers"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.Engine, userHandler *handlers.UserHandler) {
	userGroup := router.Group("/users", middlewares.AuthMiddleware())
	{
		userGroup.GET("/", userHandler.GetAllUsers)
		userGroup.GET("/:id", userHandler.GetUserById)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}
}
