package service

import "github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"

type RestaurantService interface {
	GetAllRestaurants() ([]model.RestaurantResponse, error)
	GetRestaurantById(id string) (model.RestaurantResponse, error)
	CreateRestaurant(restRequest *model.RestaurantRequest) (model.RestaurantResponse, error)
	UpdateRestaurant(id string, restRequest *model.RestaurantRequest) (model.RestaurantResponse, error)
	DeleteRestaurant(id string) error
}

type restaurantService struct{}

func NewRestaurantService() RestaurantService {
	return &restaurantService{}
}

func (r *restaurantService) CreateRestaurant(restRequest *model.RestaurantRequest) (model.RestaurantResponse, error) {
	panic("unimplemented")
}

func (r *restaurantService) DeleteRestaurant(id string) error {
	panic("unimplemented")
}

func (r *restaurantService) GetAllRestaurants() ([]model.RestaurantResponse, error) {
	panic("unimplemented")
}

func (r *restaurantService) GetRestaurantById(id string) (model.RestaurantResponse, error) {
	panic("unimplemented")
}

func (r *restaurantService) UpdateRestaurant(id string, restRequest *model.RestaurantRequest) (model.RestaurantResponse, error) {
	panic("unimplemented")
}
