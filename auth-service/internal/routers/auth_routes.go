package routers

import (
	"github.com/gin-gonic/gin"
)

func AuthRegister(router *gin.Engine) {
	router.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login successful",
		})
	})
	router.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Registration successful",
		})
	})
}
