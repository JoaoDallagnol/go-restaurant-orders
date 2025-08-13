package main

import (
	"fmt"

	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/db"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/handlers"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/routers"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.Init()

	restaurantService := service.NewRestaurantService()
	dishService := service.NewDishService()

	restauntHandler := handlers.NewRestaurantHandler(restaurantService)
	dishHandler := handlers.NewDishHandler(dishService)

	server := gin.Default()
	routers.RegisterRoutes(server, restauntHandler, dishHandler)

	port := config.AppConfig.Server.Port
	server.Run(fmt.Sprintf(":%d", port))
}
