package service

import (
	"errors"
	"strconv"

	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/repository"
	"gorm.io/gorm"
)

type RestaurantService interface {
	GetAllRestaurants() ([]model.RestaurantResponse, error)
	GetRestaurantById(id string) (model.RestaurantResponse, error)
	CreateRestaurant(restRequest *model.RestaurantRequest) (model.RestaurantResponse, error)
	UpdateRestaurant(id string, restRequest *model.RestaurantRequest) (model.RestaurantResponse, error)
	DeleteRestaurant(id string) error
}

type restaurantService struct {
	restaurantRepository repository.RestaurantRepository
}

func NewRestaurantService(restaurantRepository repository.RestaurantRepository) RestaurantService {
	return &restaurantService{restaurantRepository: restaurantRepository}
}

func (r *restaurantService) GetAllRestaurants() ([]model.RestaurantResponse, error) {
	restaurantList, err := r.restaurantRepository.GetAllRestaurants()
	if err != nil {
		return []model.RestaurantResponse{}, errs.NewInternalError(err.Error())
	}

	if len(restaurantList) == 0 {
		return []model.RestaurantResponse{}, nil
	}

	return mapper.MapRestaurantListToRestaurantResponseList(&restaurantList), nil
}

func (r *restaurantService) GetRestaurantById(id string) (model.RestaurantResponse, error) {
	restaurantId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return model.RestaurantResponse{}, errs.NewInternalError(err.Error())
	}

	restaurant, err := r.restaurantRepository.GetRestaurantById(uint(restaurantId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.RestaurantResponse{}, errs.NewRestaurantNotFound(uint(restaurantId))
		}
		return model.RestaurantResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapRestaurantToRestaurantResponse(restaurant), nil
}

func (r *restaurantService) CreateRestaurant(restRequest *model.RestaurantRequest) (model.RestaurantResponse, error) {
	restaurant := mapper.MapCreateRestaurantRequestToRestaurant(restRequest)

	createdRestaurant, err := r.restaurantRepository.CreateRestaurant(&restaurant)
	if err != nil {
		return model.RestaurantResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapRestaurantToRestaurantResponse(createdRestaurant), nil

}

func (r *restaurantService) UpdateRestaurant(id string, restRequest *model.RestaurantRequest) (model.RestaurantResponse, error) {
	restaurantId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return model.RestaurantResponse{}, errs.NewInternalError(err.Error())
	}

	existingRestaurant, err := r.restaurantRepository.GetRestaurantById(uint(restaurantId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.RestaurantResponse{}, errs.NewRestaurantNotFound(uint(restaurantId))
		}

		return model.RestaurantResponse{}, errs.NewInternalError(err.Error())
	}

	existingRestaurant.Name = restRequest.Name
	existingRestaurant.Description = restRequest.Description

	updatesRestaurant, err := r.restaurantRepository.UpdateRestaurant(existingRestaurant)
	if err != nil {
		return model.RestaurantResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapRestaurantToRestaurantResponse(updatesRestaurant), nil
}

func (r *restaurantService) DeleteRestaurant(id string) error {
	restaurantId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errs.NewInternalError(err.Error())
	}

	existingRestaurant, err := r.restaurantRepository.GetRestaurantById(uint(restaurantId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewRestaurantNotFound(uint(restaurantId))
		}

		return errs.NewInternalError(err.Error())
	}

	err = r.restaurantRepository.DeleteRestaurant(existingRestaurant)
	if err != nil {
		return errs.NewInternalError(err.Error())
	}

	return nil
}
