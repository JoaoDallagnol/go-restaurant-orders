package handlers

import (
	"net/http"
	"strconv"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderItemHandler struct {
	orderItemService service.OrderItemService
}

func NewOrderItemHandler(orderItemService service.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{orderItemService: orderItemService}
}

func (h *OrderItemHandler) GetAllOrderItems(c *gin.Context) {
	response, err := h.orderItemService.GetAllOrderItems()
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

func (h *OrderItemHandler) GetOrderItemByID(c *gin.Context) {
	id := c.Param("id")
	orderId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
		return
	}

	response, err := h.orderItemService.GetOrderItemByID(uint(orderId))
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
