package service

import (
	"errors"

	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/repository"
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
}

func NewPaymentService(paymentRepository repository.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepository: paymentRepository,
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

func (p *paymentService) CreatePayment(payment *model.PaymentRequest) (model.PaymentResponse, error) {
	panic("unimplemented")
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
