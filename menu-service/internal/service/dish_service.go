package service

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/repository"
)

type DishService interface {
	GetAllDishes() ([]model.DishResponse, error)
	GetDishById(id string) (model.DishResponse, error)
	CreateDish(dishRequest *model.DishRequest) (model.DishResponse, error)
	UpdateDish(id string, dishRequest *model.DishRequest) (model.DishResponse, error)
	DeleteDish(id string) error
}

type dishService struct {
	dishRespository repository.DishRepository
}

func NewDishService(dishRespository repository.DishRepository) DishService {
	return &dishService{dishRespository: dishRespository}
}

func (d *dishService) GetAllDishes() ([]model.DishResponse, error) {
	panic("unimplemented")
}

func (d *dishService) GetDishById(id string) (model.DishResponse, error) {
	panic("unimplemented")
}

func (d *dishService) CreateDish(dishRequest *model.DishRequest) (model.DishResponse, error) {
	//TODO IMPLEMENT 1:N LOGIC and pass Restaurant Id in the request
	panic("unimplemented")
}

func (d *dishService) DeleteDish(id string) error {
	panic("unimplemented")
}

func (d *dishService) UpdateDish(id string, dishRequest *model.DishRequest) (model.DishResponse, error) {
	panic("unimplemented")
}
