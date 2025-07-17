package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, "Usuario logado")
}

func Register(c *gin.Context) {
	c.JSON(http.StatusCreated, "Registro criado")
}
