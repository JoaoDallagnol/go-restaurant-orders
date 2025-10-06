package service

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/repository"
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
	panic("unimplemented")
}

func (p *paymentService) GetPaymentById(id uint) (model.PaymentResponse, error) {
	panic("unimplemented")
}

func (p *paymentService) CreatePayment(payment *model.PaymentRequest) (model.PaymentResponse, error) {
	panic("unimplemented")
}

func (p *paymentService) DeletePayment(id uint) error {
	panic("unimplemented")
}
