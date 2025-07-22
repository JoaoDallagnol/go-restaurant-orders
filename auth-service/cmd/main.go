package main

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/db"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/handlers"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/routers"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	db.Init()

	userRepo := repository.NewUserRepository(db.GetDB())

	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	server := gin.Default()
	routers.RegisterRoutes(server, authHandler, userHandler)

	server.Run(":8080") // Start the server on port 8080
}
