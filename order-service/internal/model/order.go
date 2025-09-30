package model

import (
	"time"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/constants"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID        uint                  `gorm:"primaryKey"`
	ClientID  uint                  `gorm:"not null"`
	Total     decimal.Decimal       `gorm:"not null"`
	Status    constants.OrderStatus `gorm:"not null"`
	CreatedAt time.Time             `gorm:"not null"`
	UpdatedAt time.Time             `gorm:"not null"`

	//1:N
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type OrderRequest struct {
	ClientID   uint               `json:"clientId" binding:"required"`
	OrderItems []OrderItemRequest `json:"items" binding:"required"`
}

type OrderResponse struct {
	ID         uint                  `json:"id"`
	ClientID   uint                  `json:"clientId"`
	Total      decimal.Decimal       `json:"total"`
	Status     constants.OrderStatus `json:"status"`
	CreatedAt  string                `json:"created_at"`
	OrderItems []OrderItemResponse   `json:"items"`
}
