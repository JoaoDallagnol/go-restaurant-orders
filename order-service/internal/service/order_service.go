package service

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/constants"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/repository"
	"github.com/shopspring/decimal"
)

type OrderService interface {
	GetAllOrders() ([]model.OrderResponse, error)
	GetOrderByID(id uint) (model.OrderResponse, error)
	CreateOrder(order *model.OrderRequest) (model.OrderResponse, error)
	UpdateOrder(id string, order *model.OrderRequest) (model.OrderResponse, error)
	DeleteOrder(id uint) error
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (o *orderService) GetAllOrders() ([]model.OrderResponse, error) {
	orderList, err := o.orderRepository.GetAllOrders()
	if err != nil {
		return nil, err
	}

	return mapper.MapOrderListToOrderResponseList(&orderList), nil
}

func (o *orderService) GetOrderByID(id uint) (model.OrderResponse, error) {
	order, err := o.orderRepository.GetOrderByID(id)
	if err != nil {
		return model.OrderResponse{}, err
	}

	return mapper.MapOrderToOrderResponse(order), nil
}

func (o *orderService) CreateOrder(order *model.OrderRequest) (model.OrderResponse, error) {
	var orderItems []model.OrderItem
	total := decimal.NewFromInt(0)

	for _, itemReq := range order.OrderItems {
		//TODO GET THE PRICE IN THE MENU-SERVICE
		price := decimal.NewFromFloat(10.0)

		orderItem := model.OrderItem{
			DishID:   itemReq.DishID,
			Quantity: itemReq.Quantity,
			Price:    price.Mul(decimal.NewFromInt(int64(itemReq.Quantity))),
		}
		total = total.Add(orderItem.Price)
		orderItems = append(orderItems, orderItem)
	}

	newOrder := &model.Order{
		ClientID:     order.ClientID,
		RestaurantID: order.RestaurantID,
		Total:        total,
		Status:       constants.StatusPending,
		OrderItems:   orderItems,
	}

	createdOrder, err := o.orderRepository.CreateOrder(newOrder)
	if err != nil {
		return model.OrderResponse{}, err
	}

	return mapper.MapOrderToOrderResponse(createdOrder), nil
}

func (o *orderService) DeleteOrder(id uint) error {
	panic("unimplemented")
}

func (o *orderService) UpdateOrder(id string, order *model.OrderRequest) (model.OrderResponse, error) {
	panic("unimplemented")
}
