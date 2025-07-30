package main

import (
	"fmt"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/db"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/handlers"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/routers"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	// Initialize DB
	db.Init()

	userRepo := repository.NewUserRepository(db.GetDB())

	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	server := gin.Default()
	routers.RegisterRoutes(server, authHandler, userHandler)

	port := config.AppConfig.Server.Port
	server.Run(fmt.Sprintf(":%d", port)) // Start the server on port 8080
}
