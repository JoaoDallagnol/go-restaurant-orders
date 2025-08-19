package service

import (
	"errors"
	"strconv"

	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/repository"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
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
	dishList, err := d.dishRespository.GetAllDishes()
	if err != nil {
		return []model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	if len(dishList) == 0 {
		return []model.DishResponse{}, nil
	}

	return mapper.MapDishListToDishResponseList(&dishList), nil
}

func (d *dishService) GetDishById(id string) (model.DishResponse, error) {
	dishId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	dish, err := d.dishRespository.GetDishById(uint(dishId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.DishResponse{}, errs.NewDishNotFound(uint(dishId))
		}
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapDishToDishResponse(dish), nil
}

func (d *dishService) CreateDish(dishRequest *model.DishRequest) (model.DishResponse, error) {
	//TODO IMPLEMENT 1:N LOGIC and pass Restaurant Id in the request
	panic("unimplemented")
}

func (d *dishService) UpdateDish(id string, dishRequest *model.DishRequest) (model.DishResponse, error) {
	dishId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	dish, err := d.dishRespository.GetDishById(uint(dishId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.DishResponse{}, errs.NewDishNotFound(uint(dishId))
		}
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	price, err := decimal.NewFromString(dishRequest.Price)
	if err != nil {
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	dish.Name = dishRequest.Name
	dish.Description = dishRequest.Description
	dish.Price = price

	updatedDish, err := d.dishRespository.UpdateDish(dish)
	if err != nil {
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapDishToDishResponse(updatedDish), nil
}

func (d *dishService) DeleteDish(id string) error {
	dishId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errs.NewInternalError(err.Error())
	}

	dish, err := d.dishRespository.GetDishById(uint(dishId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewDishNotFound(uint(dishId))
		}
		return errs.NewInternalError(err.Error())
	}

	err = d.dishRespository.DeleteDish(dish)
	if err != nil {
		return errs.NewInternalError(err.Error())
	}

	return nil
}
