package repository

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"
	"gorm.io/gorm"
)

type DishRepository interface {
	GetAllDishes() ([]model.Dish, error)
	GetDishById(id string) (*model.Dish, error)
	CreateDish(dish *model.Dish) (*model.Dish, error)
	UpdateDish(dish *model.Dish) (*model.Dish, error)
	DeleteDish(dish *model.Dish) error
}

type dishRepository struct {
	db *gorm.DB
}

func NewDishRepository(db *gorm.DB) DishRepository {
	return &dishRepository{db: db}
}

func (d *dishRepository) GetAllDishes() ([]model.Dish, error) {
	var dishies []model.Dish
	if err := d.db.Find(&dishies).Error; err != nil {
		return nil, err
	}

	return dishies, nil
}

func (d *dishRepository) GetDishById(id string) (*model.Dish, error) {
	var dish model.Dish
	if err := d.db.First(&dish, id).Error; err != nil {
		return nil, err
	}

	return &dish, nil
}

func (d *dishRepository) CreateDish(dish *model.Dish) (*model.Dish, error) {
	if err := d.db.Create(dish).Error; err != nil {
		return nil, err
	}
	return dish, nil
}

func (d *dishRepository) DeleteDish(dish *model.Dish) error {
	return d.db.Delete(&dish).Error
}

func (d *dishRepository) UpdateDish(dish *model.Dish) (*model.Dish, error) {
	if err := d.db.Save(dish).Error; err != nil {
		return nil, err
	}
	return dish, nil
}
