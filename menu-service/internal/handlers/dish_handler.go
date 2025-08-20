package handlers

import (
	"net/http"

	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"
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
	response, err := h.dishService.GetAllDishes()
	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusOK, response)
}

func (h *DishHandler) GetDishById(c *gin.Context) {
	id := c.Param("id")
	response, err := h.dishService.GetDishById(id)
	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusOK, response)
}

func (h *DishHandler) CreateDish(c *gin.Context) {
	var dishReq model.DishRequest
	if err := c.ShouldBindJSON(&dishReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.dishService.CreateDish(&dishReq)
	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusCreated, response)
}

func (h *DishHandler) UpdateDish(c *gin.Context) {
	id := c.Param("id")
	var dishReq model.DishUpdateRequest
	if err := c.ShouldBindJSON(&dishReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.dishService.UpdateDish(id, &dishReq)
	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusOK, response)
}

func (h *DishHandler) DeleteDish(c *gin.Context) {
	id := c.Param("id")
	err := h.dishService.DeleteDish(id)
	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusNoContent, model.DishResponse{})
}
