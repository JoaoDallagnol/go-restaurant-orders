package routers

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	restaurantHandler *handlers.RestaurantHandler,
	dishHandler *handlers.DishHandler,
) {
	RestaurantRegister(router, restaurantHandler)
	DishRegister(router, dishHandler)
}
