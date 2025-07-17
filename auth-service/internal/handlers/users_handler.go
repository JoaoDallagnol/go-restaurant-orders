package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "List of users")
}

func GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, "User by Id")
}
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "User update")
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusNoContent, "User deleted")
}
