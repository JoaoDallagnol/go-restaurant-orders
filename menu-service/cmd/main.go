package main

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/handlers"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	restauntHandler := handlers.NewRestaurantHandler()
	dishHandler := handlers.NewDishHandler()

	server := gin.Default()
	routers.RegisterRoutes(server, restauntHandler, dishHandler)
	server.Run(":8080")
}
