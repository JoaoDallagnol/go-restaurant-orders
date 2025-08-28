package main

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	server := gin.Default()
	server.Run(":8082")
}
