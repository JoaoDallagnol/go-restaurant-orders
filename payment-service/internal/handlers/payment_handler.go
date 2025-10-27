package handlers

import (
	"net/http"
	"strconv"

	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/service"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService service.PaymentService
}

func NewPaymentHandler(paymentService service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) GetAllPayments(c *gin.Context) {
	response, err := h.paymentService.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *PaymentHandler) GetPaymentById(c *gin.Context) {
	id := c.Param("id")
	paymentId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
		return
	}

	response, err := h.paymentService.GetPaymentById(uint(paymentId))
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

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var paymentReq model.PaymentRequest
	if err := c.ShouldBindJSON(&paymentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.paymentService.CreatePayment(&paymentReq)
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
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *PaymentHandler) DeletePayment(c *gin.Context) {
	id := c.Param("id")
	paymentId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   errs.CodeInternalError,
			"details": err.Error(),
		})
		return
	}

	err = h.paymentService.DeletePayment(uint(paymentId))
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
		return
	}

	c.JSON(http.StatusNoContent, model.PaymentRequest{})
}
