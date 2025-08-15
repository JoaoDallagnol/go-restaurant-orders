package model

import "time"

type Restaurant struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	//1:N
	Dishes []Dish `gorm:"foreignKey:RestaurantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type RestaurantRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type RestaurantResponse struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   string         `json:"created_at"`
	Dishes      []DishResponse `json:"dishes"`
}
