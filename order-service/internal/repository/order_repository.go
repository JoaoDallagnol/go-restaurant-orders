package repository

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders() ([]model.Order, error)
	GetOrderByID(id uint) (*model.Order, error)
	CreateOrder(order *model.Order) (*model.Order, error)
	UpdateOrder(order *model.Order) error
	DeleteOrder(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (o *orderRepository) GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	if err := o.db.Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderRepository) GetOrderByID(id uint) (*model.Order, error) {
	var order model.Order
	if err := o.db.Preload("Items").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *orderRepository) CreateOrder(order *model.Order) (*model.Order, error) {
	if err := o.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (o *orderRepository) DeleteOrder(id uint) error {
	if err := o.db.Delete(&model.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) UpdateOrder(order *model.Order) error {
	if err := o.db.Save(order).Error; err != nil {
		return err
	}
	return nil
}
