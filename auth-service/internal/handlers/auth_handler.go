package handlers

import (
	"net/http"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

var authService = service.NewAuthService()

func Login(c *gin.Context) {
	var loginReq model.UserLoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := authService.Login(loginReq)
	c.JSON(http.StatusOK, response)
}

func Register(c *gin.Context) {
	var userRequest model.RegisterUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {

		// MAKE EROR TREATMENT
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := authService.RegisterUser(userRequest)
	c.JSON(http.StatusCreated, response)
}
