package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", handlers.GetAllUsers)
		userGroup.GET("/:id", handlers.GetUserById)
		userGroup.PUT("/:id", handlers.UpdateUser)
		userGroup.DELETE("/:id", handlers.DeleteUser)
	}
}
