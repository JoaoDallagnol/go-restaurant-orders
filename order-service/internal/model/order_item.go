package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID        uint            `gorm:"primaryKey"`
	DishID    uint            `gorm:"not null"`
	Quantity  int             `gorm:"not null"`
	Price     decimal.Decimal `gorm:"not null"`
	CreatedAt time.Time       `gorm:"not null"`
	UpdatedAt time.Time       `gorm:"not null"`

	// Foreign key
	OrderID uint  `gorm:"not null;index"`
	Order   Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type OrderItemRequest struct {
	DishID   uint `json:"dishId" binding:"required"`
	Quantity int  `json:"quantity" binding:"required"`
}

type OrderItemResponse struct {
	ID        uint            `json:"id"`
	OrderID   uint            `json:"orderId"`
	DishID    uint            `json:"dishId"`
	Quantity  int             `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
	CreatedAt string          `json:"created_at"`
}
