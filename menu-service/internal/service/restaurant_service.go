package service

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/repository"
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
	panic("unimplemented")
}

func (r *restaurantService) CreateRestaurant(restRequest *model.RestaurantRequest) (model.RestaurantResponse, error) {
	panic("unimplemented")
}

func (r *restaurantService) DeleteRestaurant(id string) error {
	panic("unimplemented")
}

func (r *restaurantService) UpdateRestaurant(id string, restRequest *model.RestaurantRequest) (model.RestaurantResponse, error) {
	panic("unimplemented")
}
