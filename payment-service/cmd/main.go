package main

import (
	"fmt"

	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/db"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/client"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/handlers"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/publishers"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/repository"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/routers"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.Init()

	orderPaymentPublisher, _ := publishers.NewOrderServicePublisher(config.AppConfig)
	orderClient := client.NewOrderClient(config.AppConfig)
	paymentRepository := repository.NewPaymentRepository(db.DB)
	paymentService := service.NewPaymentService(paymentRepository, orderClient, orderPaymentPublisher)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	server := gin.Default()
	routers.RegisterRoutes(server, paymentHandler)

	port := config.AppConfig.Server.Port
	server.Run(fmt.Sprintf(":%d", port))
}
