package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItems struct {
	ID        uint            `gorm:"primaryKey"`
	OrderId   uint            `gorm:"not null"`
	DishId    uint            `gorm:"not null"`
	Quantity  int             `gorm:"not null"`
	Price     decimal.Decimal `gorm:"not null"`
	CreatedAt time.Time       `gorm:"not null"`
	UpdatedAt time.Time       `gorm:"not null"`

	// Foreign key
	OrderID uint  `gorm:"not null;index"`
	Order   Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
