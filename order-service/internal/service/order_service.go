package service

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/repository"
)

type OrderService interface {
	GetAllOrders() ([]model.OrderResponse, error)
	GetOrderByID(id uint) (*model.OrderResponse, error)
	CreateOrder(order *model.OrderRequest) (*model.OrderResponse, error)
	UpdateOrder(id string, order *model.OrderRequest) (*model.OrderResponse, error)
	DeleteOrder(id uint) error
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (o *orderService) GetAllOrders() ([]model.OrderResponse, error) {
	panic("unimplemented")
}

func (o *orderService) GetOrderByID(id uint) (*model.OrderResponse, error) {
	panic("unimplemented")
}

func (o *orderService) CreateOrder(order *model.OrderRequest) (*model.OrderResponse, error) {
	panic("unimplemented")
}

func (o *orderService) DeleteOrder(id uint) error {
	panic("unimplemented")
}

func (o *orderService) UpdateOrder(id string, order *model.OrderRequest) (*model.OrderResponse, error) {
	panic("unimplemented")
}
