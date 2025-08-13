package repository

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"
	"gorm.io/gorm"
)

type RestaurantRepository interface {
	GetAllRestaurants() ([]model.Restaurant, error)
	GetRestaurantById(id uint) (*model.Restaurant, error)
	CreateRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error)
	UpdateRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error)
	DeleteRestaurant(restaurant *model.Restaurant) error
}

type restaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return &restaurantRepository{db: db}
}

func (r *restaurantRepository) GetAllRestaurants() ([]model.Restaurant, error) {
	var restaurants []model.Restaurant
	if err := r.db.Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (r *restaurantRepository) GetRestaurantById(id uint) (*model.Restaurant, error) {
	var restaurant model.Restaurant
	if err := r.db.First(&restaurant, id).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func (r *restaurantRepository) CreateRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error) {
	if err := r.db.Create(restaurant).Error; err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (r *restaurantRepository) DeleteRestaurant(restaurant *model.Restaurant) error {
	return r.db.Delete(&restaurant).Error
}

func (r *restaurantRepository) UpdateRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error) {
	if err := r.db.Save(restaurant).Error; err != nil {
		return nil, err
	}
	return restaurant, nil
}
