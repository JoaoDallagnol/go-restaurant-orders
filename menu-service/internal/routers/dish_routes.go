package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func DishRegister(router *gin.Engine, dishHandler *handlers.DishHandler) {
	dishGroup := router.Group("/dishes")
	{
		dishGroup.GET("/", dishHandler.GetAllDishes)
		dishGroup.GET("/:id", dishHandler.GetDishById)
		dishGroup.POST("/", dishHandler.CreateDish)
		dishGroup.PUT("/:id", dishHandler.UpdateDish)
		dishGroup.DELETE("/:id", dishHandler.DeleteDish)
	}
}
