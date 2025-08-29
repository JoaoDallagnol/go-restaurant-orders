package main

import (
	"fmt"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	server := gin.Default()

	port := config.AppConfig.Server.Port
	server.Run(fmt.Sprintf(":%d", port))
}
