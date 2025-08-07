package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Dish struct {
	ID          uint            `gorm:"primaryKey"`
	Name        string          `gorm:"not null"`
	Description string          `gorm:"not null"`
	Price       decimal.Decimal `gorm:"not null"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"`
}

type DishRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       string `json:"price" binding:"required"`
}

type DishResponse struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       string `json:"price" binding:"required"`
	CreatedAt   string `json:"created_at"`
}
