package handlers

import (
	"net/http"

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
	c.JSON(http.StatusOK, "Restaurant List")
}

func (h *RestaurantHandler) GetRestaurantById(c *gin.Context) {
	c.JSON(http.StatusOK, "Restaurant by Id")
}

func (h *RestaurantHandler) CreateRestaurant(c *gin.Context) {
	c.JSON(http.StatusCreated, "Restaurant Created")
}

func (h *RestaurantHandler) UpdateRestaurant(c *gin.Context) {
	c.JSON(http.StatusOK, "Restaurant Updated")
}

func (h *RestaurantHandler) DeleteRestaurant(c *gin.Context) {
	c.JSON(http.StatusNoContent, "Restaurant Deleted")
}
