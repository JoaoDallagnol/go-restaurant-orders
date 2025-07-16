package main

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routers.RegisterRoutes(server)

	server.Run(":8080") // Start the server on port 8080
}
