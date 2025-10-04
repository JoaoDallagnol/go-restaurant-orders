package model

import "time"

type Payment struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `gorm:"not null"`
	Amount    string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type PaymentRequest struct {
	OrderId uint   `json:"orderId" binding:"required"`
	Amount  string `json:"amount" binding:"required"`
}

type PaymentResponse struct {
	ID        uint      `json:"id"`
	OrderID   uint      `json:"orderId"`
	Amount    string    `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
