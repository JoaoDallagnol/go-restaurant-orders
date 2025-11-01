package service

import (
	"errors"
	"log"

	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/client"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/constants"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/publishers"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/repository"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PaymentService interface {
	GetAllPayments() ([]model.PaymentResponse, error)
	GetPaymentById(id uint) (model.PaymentResponse, error)
	CreatePayment(payment *model.PaymentRequest) (model.PaymentResponse, error)
	DeletePayment(id uint) error
}

type paymentService struct {
	paymentRepository repository.PaymentRepository
	orderClient       client.OrderClient
	publisher         *publishers.OrderServicePublisher
}

func NewPaymentService(
	paymentRepository repository.PaymentRepository,
	orderClient client.OrderClient,
	publisher *publishers.OrderServicePublisher,
) PaymentService {
	return &paymentService{
		paymentRepository: paymentRepository,
		orderClient:       orderClient,
		publisher:         publisher,
	}
}

func (p *paymentService) GetAllPayments() ([]model.PaymentResponse, error) {
	paymentList, err := p.paymentRepository.GetAllPayments()
	if err != nil {
		return nil, err
	}

	return mapper.MapPaymentListToPaymentResponseList(&paymentList), nil
}

func (p *paymentService) GetPaymentById(id uint) (model.PaymentResponse, error) {
	payment, err := p.paymentRepository.GetPaymentById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.PaymentResponse{}, errs.NewPaymentNotFound(id)
		}
		return model.PaymentResponse{}, errs.NewInternalError(err.Error())
	}
	return mapper.MapPaymentToPaymentResponse(payment), nil
}

func (p *paymentService) CreatePayment(paymentReq *model.PaymentRequest) (model.PaymentResponse, error) {
	order, err := p.orderClient.GetOrderById(paymentReq.OrderId)
	if err != nil {
		return model.PaymentResponse{}, errs.NewOrderServiceIntegrationError()
	}

	var status constants.PaymentStatus
	amountDecimal, err := decimal.NewFromString(paymentReq.Amount)
	if err != nil {
		status = constants.CANCELLED
	} else {
		if order.Total.GreaterThan(amountDecimal) {
			status = constants.DECLINED
		} else {
			status = constants.APPROVED
		}
	}

	payment := mapper.MapPaymentRequestToPayment(paymentReq, status)
	createdPayment, err := p.paymentRepository.CreatePayment(&payment)
	if err != nil {
		return model.PaymentResponse{}, errs.NewInternalError(err.Error())
	}

	var orderStatus constants.OrderStatus
	switch status {
	case constants.APPROVED:
		orderStatus = constants.StatusConfirmed
	case constants.CANCELLED, constants.DECLINED:
		orderStatus = constants.StatusCancelled
	default:
		orderStatus = constants.StatusPending
	}

	if err := p.publisher.Publish(paymentReq.OrderId, orderStatus); err != nil {
		log.Printf("⚠️ Failed to publish order status update: %v", err)
	}

	return mapper.MapPaymentToPaymentResponse(createdPayment), nil
}

func (p *paymentService) DeletePayment(id uint) error {
	payment, err := p.paymentRepository.GetPaymentById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewPaymentNotFound(id)
		}
		return errs.NewInternalError(err.Error())
	}

	err = p.paymentRepository.DeletePayment(payment)
	if err != nil {
		return errs.NewInternalError(err.Error())
	}

	return nil
}
