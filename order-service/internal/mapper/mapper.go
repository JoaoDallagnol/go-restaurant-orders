package mapper

import "github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/model"

func MapOrderItemToOrderItemResponse(orderItem *model.OrderItem) model.OrderItemResponse {
	return model.OrderItemResponse{
		ID:        orderItem.ID,
		OrderID:   orderItem.OrderID,
		DishID:    orderItem.DishID,
		Quantity:  orderItem.Quantity,
		CreatedAt: orderItem.CreatedAt.String(),
	}
}

func MapOrderItemListToOrderItemResponseList(orderItemList *[]model.OrderItem) []model.OrderItemResponse {
	orderItemResponseList := make([]model.OrderItemResponse, len(*orderItemList))

	for i, orderItem := range *orderItemList {
		orderItemResponseList[i] = MapOrderItemToOrderItemResponse(&orderItem)
	}
	return orderItemResponseList
}
