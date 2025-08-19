package mapper

import "github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"

func MapRestaurantToRestaurantResponse(restaurant *model.Restaurant) model.RestaurantResponse {
	dishes := make([]model.DishResponse, 0, len(restaurant.Dishes))
	for _, dish := range restaurant.Dishes {
		dishes = append(dishes, model.DishResponse{
			ID:          dish.ID,
			Name:        dish.Name,
			Description: dish.Description,
			Price:       dish.Price.String(),
			CreatedAt:   dish.CreatedAt.String(),
		})
	}

	return model.RestaurantResponse{
		ID:          restaurant.ID,
		Name:        restaurant.Name,
		Description: restaurant.Description,
		CreatedAt:   restaurant.CreatedAt.String(),
		Dishes:      dishes,
	}
}

func MapRestaurantListToRestaurantResponseList(restaurantList *[]model.Restaurant) []model.RestaurantResponse {
	response := make([]model.RestaurantResponse, 0, len(*restaurantList))

	for _, restaurant := range *restaurantList {
		response = append(response, MapRestaurantToRestaurantResponse(&restaurant))
	}

	return response
}

func MapCreateRestaurantRequestToRestaurant(restReq *model.RestaurantRequest) model.Restaurant {
	return model.Restaurant{
		Name:        restReq.Name,
		Description: restReq.Description,
	}
}

func MapDishToDishResponse(dish *model.Dish) model.DishResponse {
	return model.DishResponse{
		ID:           dish.ID,
		Name:         dish.Name,
		Description:  dish.Description,
		Price:        dish.Price.String(),
		RestaurantID: dish.RestaurantID,
		CreatedAt:    dish.CreatedAt.String(),
	}
}

func MapDishListToDishResponseList(dishList *[]model.Dish) []model.DishResponse {
	response := make([]model.DishResponse, 0, len(*dishList))

	for _, dish := range *dishList {
		response = append(response, MapDishToDishResponse(&dish))
	}
	return response
}
