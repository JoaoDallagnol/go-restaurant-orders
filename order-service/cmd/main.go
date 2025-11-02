package main

import (
	"fmt"
	"log"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/db"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/client"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/consumers"
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
	authClient := client.NewAuthClient(config.AppConfig)

	orderRepository := repository.NewOrderRepository(db.DB)
	orderItemRepository := repository.NewOrderItemRepository(db.DB)

	orderService := service.NewOrderService(orderRepository, orderItemRepository, menuClient, authClient)
	orderItemService := service.NewOrderItemService(orderItemRepository)

	orderHandler := handlers.NewOrderHandler(orderService)
	orderItemHandler := handlers.NewOrderItemHandler(orderItemService)

	consumer, err := consumers.NewOrderStatusConsumer(config.AppConfig, orderRepository)
	if err != nil {
		log.Fatalf("❌ Failed to initialize RabbitMQ consumer: %v", err)
	}

	go func() {
		if err := consumer.Start(); err != nil {
			log.Fatalf("❌ Consumer crashed: %v", err)
		}
	}()

	server := gin.Default()
	routers.RegisterRoutes(server, orderHandler, orderItemHandler)

	port := config.AppConfig.Server.Port
	server.Run(fmt.Sprintf(":%d", port))
}
