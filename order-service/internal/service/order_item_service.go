package service

import (
	"errors"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/repository"
	"gorm.io/gorm"
)

type OrderItemService interface {
	GetAllOrderItems() ([]model.OrderItemResponse, error)
	GetOrderItemByID(id uint) (model.OrderItemResponse, error)
}

type orderItemService struct {
	orderItemRepository repository.OrderItemRepository
}

func NewOrderItemService(orderItemRepository repository.OrderItemRepository) OrderItemService {
	return &orderItemService{orderItemRepository: orderItemRepository}
}

func (o *orderItemService) GetAllOrderItems() ([]model.OrderItemResponse, error) {
	orderItemList, err := o.orderItemRepository.GetAllOrderItems()
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	return mapper.MapOrderItemListToOrderItemResponseList(&orderItemList), nil
}

func (o *orderItemService) GetOrderItemByID(id uint) (model.OrderItemResponse, error) {
	orderItem, err := o.orderItemRepository.GetOrderItemByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.OrderItemResponse{}, errs.NewOrderItemNotFound(id)
		}
		return model.OrderItemResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapOrderItemToOrderItemResponse(orderItem), nil
}
