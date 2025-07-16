package routers

import (
	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "List of users",
			})
		})

		userGroup.GET("/:id", func(c *gin.Context) {
			userID := c.Param("id")
			c.JSON(200, gin.H{
				"message": "User details",
				"id":      userID,
			})
		})
		userGroup.PUT("/:id", func(c *gin.Context) {
			userID := c.Param("id")
			c.JSON(200, gin.H{
				"message": "User updated",
				"id":      userID,
			})
		})
		userGroup.DELETE("/:id", func(c *gin.Context) {
			userID := c.Param("id")
			c.JSON(200, gin.H{
				"message": "User deleted",
				"id":      userID,
			})
		})
	}
}
