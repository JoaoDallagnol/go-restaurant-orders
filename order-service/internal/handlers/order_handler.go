package handlers

import (
	"net/http"
	"strconv"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	response, err := h.orderService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	orderId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
	}

	response, err := h.orderService.GetOrderByID(uint(orderId))
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

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var orderReq model.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.orderService.CreateOrder(&orderReq)
	if err != nil {
		if apiErr, ok := err.(errs.CodedError); ok {
			c.JSON(errs.MapErrorCodeToStatus(apiErr.GetCode()), gin.H{
				"error":   apiErr.GetCode(),
				"details": apiErr.GetDetails(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, response)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	orderId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
	}

	var orderReq model.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.orderService.UpdateOrder(uint(orderId), &orderReq)
	if err != nil {
		if apiErr, ok := err.(errs.CodedError); ok {
			c.JSON(errs.MapErrorCodeToStatus(apiErr.GetCode()), gin.H{
				"error":   apiErr.GetCode(),
				"details": apiErr.GetDetails(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	orderId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
	}

	err = h.orderService.DeleteOrder(uint(orderId))
	if err != nil {
		if apiErr, ok := err.(errs.CodedError); ok {
			c.JSON(errs.MapErrorCodeToStatus(apiErr.GetCode()), gin.H{
				"error":   apiErr.GetCode(),
				"details": apiErr.GetDetails(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
	}

	c.JSON(http.StatusNoContent, model.OrderResponse{})
}
