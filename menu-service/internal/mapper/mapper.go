package mapper

import "github.com/JoaoDallagnol/go-restaurant-orders/menu-service/internal/model"

func MapRestaurantToRestaurantResponse(restaurant *model.Restaurant) model.RestaurantResponse {
	return model.RestaurantResponse{
		ID:          restaurant.ID,
		Name:        restaurant.Name,
		Description: restaurant.Description,
		CreatedAt:   restaurant.CreatedAt.String(),
	}
}

func MapRestaurantListToRestaurantResponseList(restaurantList *[]model.Restaurant) []model.RestaurantResponse {
	response := make([]model.RestaurantResponse, 0, len(*restaurantList))

	for _, restaurant := range *restaurantList {
		response = append(response, MapRestaurantToRestaurantResponse(&restaurant))
	}

	return response
}
