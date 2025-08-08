package handlers

import (
	"net/http"

	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/service"
	"github.com/gin-gonic/gin"
)

type DishHandler struct {
	dishService service.DishService
}

func NewDishHandler(dishService service.DishService) *DishHandler {
	return &DishHandler{dishService: dishService}
}

func (h *DishHandler) GetAllDishes(c *gin.Context) {
	c.JSON(http.StatusOK, "Dishes List")
}

func (h *DishHandler) GetDishById(c *gin.Context) {
	c.JSON(http.StatusOK, "Dish by Id")
}

func (h *DishHandler) CreateDish(c *gin.Context) {
	c.JSON(http.StatusCreated, "Dish Ureated")
}

func (h *DishHandler) UpdateDish(c *gin.Context) {
	c.JSON(http.StatusOK, "Dish Updated")
}

func (h *DishHandler) DeleteDish(c *gin.Context) {
	c.JSON(http.StatusNoContent, "Dish Deleted")
}
