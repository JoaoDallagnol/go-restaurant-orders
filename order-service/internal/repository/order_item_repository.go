package repository

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/model"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	GetAllOrderItems() ([]model.OrderItem, error)
	GetOrderItemByID(id uint) (*model.OrderItem, error)
	DeleteOrderItemsByOrderID(ordemItem *model.OrderItem) error
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{db: db}
}

func (o *orderItemRepository) GetAllOrderItems() ([]model.OrderItem, error) {
	var orderItems []model.OrderItem
	if err := o.db.Find(&orderItems).Error; err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (o *orderItemRepository) GetOrderItemByID(id uint) (*model.OrderItem, error) {
	var orderItem model.OrderItem
	if err := o.db.First(&orderItem, id).Error; err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (o *orderItemRepository) DeleteOrderItemsByOrderID(orderItem *model.OrderItem) error {
	return o.db.Delete(&orderItem).Error
}
