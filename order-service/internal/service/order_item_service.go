package service

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/repository"
)

type OrderItemService interface {
	GetAllOrderItems() ([]model.OrderItemResponse, error)
	GetOrderItemByID(id uint) (*model.OrderItemResponse, error)
}

type orderItemService struct {
	orderItemRepository repository.OrderItemRepository
}

func NewOrderItemService(orderItemRepository repository.OrderItemRepository) OrderItemService {
	return &orderItemService{orderItemRepository: orderItemRepository}
}

func (o *orderItemService) GetAllOrderItems() ([]model.OrderItemResponse, error) {
	panic("unimplemented")
}

func (o *orderItemService) GetOrderItemByID(id uint) (*model.OrderItemResponse, error) {
	panic("unimplemented")
}
