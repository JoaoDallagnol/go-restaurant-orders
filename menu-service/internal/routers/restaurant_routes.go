package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RestaurantRegister(router *gin.Engine, restaurantHandler *handlers.RestaurantHandler) {
	restaurantGroup := router.Group("/restaurants")
	{
		restaurantGroup.GET("/", restaurantHandler.GetAllRestaurants)
		restaurantGroup.GET("/:id", restaurantHandler.GetRestaurantById)
		restaurantGroup.POST("/", restaurantHandler.CreateRestaurant)
		restaurantGroup.PUT("/:id", restaurantHandler.UpdateRestaurant)
		restaurantGroup.DELETE("/:id", restaurantHandler.DeleteRestaurant)
	}
}
