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
	UpdateDish(id string, dishRequest *model.DishUpdateRequest) (model.DishResponse, error)
	DeleteDish(id string) error
}

type dishService struct {
	dishRespository      repository.DishRepository
	restaurantRepository repository.RestaurantRepository
}

func NewDishService(dishRespository repository.DishRepository, restaurantRepository repository.RestaurantRepository) DishService {
	return &dishService{
		dishRespository:      dishRespository,
		restaurantRepository: restaurantRepository,
	}
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
	restaurantId, err := strconv.ParseUint(dishRequest.RestaurantId, 10, 32)
	if err != nil {
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	if _, err := d.restaurantRepository.GetRestaurantById(uint(restaurantId)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.DishResponse{}, errs.NewRestaurantNotFound(uint(restaurantId))
		}
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	dish := mapper.MapDishRequestToDish(dishRequest)
	if err != nil {
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	dish.RestaurantID = uint(restaurantId)

	createdDish, err := d.dishRespository.CreateDish(&dish)
	if err != nil {
		return model.DishResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapDishToDishResponse(createdDish), nil
}

func (d *dishService) UpdateDish(id string, dishRequest *model.DishUpdateRequest) (model.DishResponse, error) {
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
