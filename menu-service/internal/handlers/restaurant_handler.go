package handlers

import (
	"net/http"

	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/service"
	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	restaurantService service.RestaurantService
}

func NewRestaurantHandler(restaurantService service.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{restaurantService: restaurantService}
}

func (h *RestaurantHandler) GetAllRestaurants(c *gin.Context) {
	response, err := h.restaurantService.GetAllRestaurants()
	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusOK, response)
}

func (h *RestaurantHandler) GetRestaurantById(c *gin.Context) {
	id := c.Param("id")
	response, err := h.restaurantService.GetRestaurantById(id)
	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusOK, response)
}

func (h *RestaurantHandler) CreateRestaurant(c *gin.Context) {
	var restaurantReq model.RestaurantRequest
	if err := c.ShouldBindJSON(&restaurantReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.restaurantService.CreateRestaurant(&restaurantReq)
	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusCreated, response)
}

func (h *RestaurantHandler) UpdateRestaurant(c *gin.Context) {
	id := c.Param("id")
	var restaurantReq model.RestaurantRequest

	if err := c.ShouldBindJSON(&restaurantReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.restaurantService.UpdateRestaurant(id, &restaurantReq)
	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusOK, response)
}

func (h *RestaurantHandler) DeleteRestaurant(c *gin.Context) {
	id := c.Param("id")
	err := h.restaurantService.DeleteRestaurant(id)

	if err != nil {
		panic("error: " + err.Error())
	}

	c.JSON(http.StatusNoContent, model.DishResponse{})
}
