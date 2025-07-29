package handlers

import (
	"net/http"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/errs"
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
	response, err := h.userService.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	response, err := h.userService.GetUserById(id)

	if err != nil {
		if apiErr, ok := err.(errs.CodedError); ok {
			status := errs.MapErrorCodeToStatus(apiErr.GetCode())
			c.JSON(status, gin.H{
				"error":   apiErr.GetCode(),
				"details": apiErr.GetDetails(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
		return
	}
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
