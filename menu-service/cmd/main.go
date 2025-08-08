package main

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/handlers"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/routers"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	restaurantService := service.NewRestaurantService()
	dishService := service.NewDishService()

	restauntHandler := handlers.NewRestaurantHandler(restaurantService)
	dishHandler := handlers.NewDishHandler(dishService)

	server := gin.Default()
	routers.RegisterRoutes(server, restauntHandler, dishHandler)
	server.Run(":8080")
}
