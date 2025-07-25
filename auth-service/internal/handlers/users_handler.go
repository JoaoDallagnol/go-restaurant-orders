package handlers

import (
	"net/http"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	response := h.userService.GetAllUser()
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	response := h.userService.GetUserById(id)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userReq model.RegisterUserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := h.userService.UpdateUser(id, &userReq)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	h.userService.DeleteUser(id)
	c.JSON(http.StatusNoContent, model.User{})
}
