package model

import (
	"time"

	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/constants"
)

type Payment struct {
	ID        uint                    `gorm:"primaryKey"`
	OrderID   uint                    `gorm:"not null"`
	Amount    string                  `gorm:"not null"`
	Status    constants.PaymentStatus `gorm:"not null"`
	CreatedAt time.Time               `gorm:"autoCreateTime"`
}

type PaymentRequest struct {
	OrderId uint   `json:"orderId" binding:"required"`
	Amount  string `json:"amount" binding:"required"`
}

type PaymentResponse struct {
	ID        uint                    `json:"id"`
	OrderID   uint                    `json:"orderId"`
	Amount    string                  `json:"amount"`
	Status    constants.PaymentStatus `json:"status"`
	CreatedAt string                  `json:"createdAt"`
}
