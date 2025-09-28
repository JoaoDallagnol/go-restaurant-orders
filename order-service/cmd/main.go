package main

import (
	"fmt"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/db"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/client"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/handlers"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/repository"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/routers"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.Init()

	menuClient := client.NewMenuClient(config.AppConfig)

	orderRespository := repository.NewOrderRepository(db.DB)
	orderItemRepository := repository.NewOrderItemRepository(db.DB)

	orderService := service.NewOrderService(orderRespository, menuClient)
	orderItemService := service.NewOrderItemService(orderItemRepository)

	orderHandler := handlers.NewOrderHandler(orderService)
	orderItemHandler := handlers.NewOrderItemHandler(orderItemService)

	server := gin.Default()
	routers.RegisterRoutes(server, orderHandler, orderItemHandler)

	port := config.AppConfig.Server.Port
	server.Run(fmt.Sprintf(":%d", port))
}
