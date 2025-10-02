package service

import (
	"errors"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/client"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/constants"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/repository"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderService interface {
	GetAllOrders() ([]model.OrderResponse, error)
	GetOrderByID(id uint) (model.OrderResponse, error)
	CreateOrder(order *model.OrderRequest) (model.OrderResponse, error)
	UpdateOrder(id uint, order *model.OrderRequest) (model.OrderResponse, error)
	DeleteOrder(id uint) error
}

type orderService struct {
	orderRepository repository.OrderRepository
	menuClient      client.MenuClient
	authClient      client.AuthClient
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	menuClient client.MenuClient,
	authClient client.AuthClient,
) OrderService {
	return &orderService{
		orderRepository: orderRepository,
		menuClient:      menuClient,
		authClient:      authClient,
	}
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.OrderResponse{}, errs.NewOrderNotFound(id)
		}
		return model.OrderResponse{}, errs.NewInternalError(err.Error())
	}
	return mapper.MapOrderToOrderResponse(order), nil
}

func (o *orderService) CreateOrder(order *model.OrderRequest) (model.OrderResponse, error) {
	var orderItems []model.OrderItem
	total := decimal.NewFromInt(0)

	_, err := o.authClient.GetUserById(order.ClientID)
	if err != nil {
		return model.OrderResponse{}, errs.NewAuthServiceIntegrationError()
	}

	for _, itemReq := range order.OrderItems {
		dish, err := o.menuClient.GetDishByID(itemReq.DishID)
		if err != nil {
			return model.OrderResponse{}, errs.NewMenuServiceIntegrationError()
		}

		price, err := decimal.NewFromString(dish.Price)
		if err != nil {
			return model.OrderResponse{}, errs.NewInternalError(err.Error())
		}

		orderItem := model.OrderItem{
			DishID:       itemReq.DishID,
			RestaurantID: dish.RestaurantID,
			Quantity:     itemReq.Quantity,
			Price:        price.Mul(decimal.NewFromInt(int64(itemReq.Quantity))),
		}
		total = total.Add(orderItem.Price)
		orderItems = append(orderItems, orderItem)
	}

	newOrder := &model.Order{
		ClientID:   order.ClientID,
		Total:      total,
		Status:     constants.StatusPending,
		OrderItems: orderItems,
	}

	createdOrder, err := o.orderRepository.CreateOrder(newOrder)
	if err != nil {
		return model.OrderResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapOrderToOrderResponse(createdOrder), nil
}

func (o *orderService) UpdateOrder(id uint, order *model.OrderRequest) (model.OrderResponse, error) {
	existingOrder, err := o.orderRepository.GetOrderByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.OrderResponse{}, errs.NewOrderNotFound(id)
		}
		return model.OrderResponse{}, errs.NewInternalError(err.Error())
	}

	_, err = o.authClient.GetUserById(order.ClientID)
	if err != nil {
		return model.OrderResponse{}, errs.NewAuthServiceIntegrationError()
	}

	existingOrder.ClientID = order.ClientID

	var updatedItems []model.OrderItem
	total := decimal.NewFromInt(0)

	for _, itemReq := range order.OrderItems {
		dish, err := o.menuClient.GetDishByID(itemReq.DishID)
		if err != nil {
			return model.OrderResponse{}, errs.NewMenuServiceIntegrationError()
		}

		price, err := decimal.NewFromString(dish.Price)
		if err != nil {
			return model.OrderResponse{}, errs.NewInternalError(err.Error())
		}

		orderItem := model.OrderItem{
			DishID:       itemReq.DishID,
			RestaurantID: dish.RestaurantID,
			Quantity:     itemReq.Quantity,
			Price:        price.Mul(decimal.NewFromInt(int64(itemReq.Quantity))),
			OrderID:      existingOrder.ID,
		}
		total = total.Add(orderItem.Price)
		updatedItems = append(updatedItems, orderItem)
	}

	existingOrder.OrderItems = updatedItems
	existingOrder.Total = total

	if err := o.orderRepository.UpdateOrder(existingOrder); err != nil {
		return model.OrderResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapOrderToOrderResponse(existingOrder), nil
}

func (o *orderService) DeleteOrder(id uint) error {

	_, err := o.orderRepository.GetOrderByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewOrderNotFound(id)
		}

		return errs.NewInternalError(err.Error())
	}

	err = o.orderRepository.DeleteOrder(id)
	if err != nil {
		return errs.NewInternalError(err.Error())
	}

	return nil
}
