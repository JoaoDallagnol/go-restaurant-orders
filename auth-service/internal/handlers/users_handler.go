package handlers

import (
	"net/http"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

var userService = service.NewUserService()

func GetAllUsers(c *gin.Context) {
	response := userService.GetAllUser()
	c.JSON(http.StatusOK, response)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	response := userService.GetUserById(id)
	c.JSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userReq model.RegisterUserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := userService.UpdateUser(id, userReq)
	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	response := userService.DeleteUser(id)
	c.JSON(http.StatusNoContent, response)
}
