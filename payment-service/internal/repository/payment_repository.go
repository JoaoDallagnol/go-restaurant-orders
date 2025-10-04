package repository

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/model"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetAllPayments() ([]model.Payment, error)
	GetPaymentById(id uint) (*model.Payment, error)
	CreatePayment(payment *model.Payment) (*model.Payment, error)
	UpdatePayment(payment *model.Payment) (*model.Payment, error)
	DeletePayment(payment *model.Payment) error
}

type paymentRespository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRespository{db: db}
}

func (p *paymentRespository) GetAllPayments() ([]model.Payment, error) {
	var payments []model.Payment
	if err := p.db.Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

func (p *paymentRespository) GetPaymentById(id uint) (*model.Payment, error) {
	var payment model.Payment
	if err := p.db.First(&payment, id).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

func (p *paymentRespository) CreatePayment(payment *model.Payment) (*model.Payment, error) {
	if err := p.db.Create(payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func (p *paymentRespository) UpdatePayment(payment *model.Payment) (*model.Payment, error) {
	if err := p.db.Save(payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func (p *paymentRespository) DeletePayment(payment *model.Payment) error {
	return p.db.Delete(&payment).Error
}
