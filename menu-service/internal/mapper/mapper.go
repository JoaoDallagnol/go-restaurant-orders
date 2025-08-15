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
