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

	// Foreign key
	RestaurantID uint       `gorm:"not null;index"`
	Restaurant   Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type DishRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       string `json:"price" binding:"required"`
}

type DishResponse struct {
	ID           uint   `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Price        string `json:"price" binding:"required"`
	CreatedAt    string `json:"created_at"`
	RestaurantID uint   `json:"restaurant_id,omitempty"`
}
