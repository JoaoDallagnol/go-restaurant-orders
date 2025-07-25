package mapper

import "github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"

func MapCreateUserRequestToUser(dto *model.RegisterUserRequest) model.User {
	return model.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func MapUserToUserResponse(user *model.User) model.UserResponse {
	return model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}
}

func MapUserListToUserResponseList(userList *[]model.User) []model.UserResponse {
	response := make([]model.UserResponse, 0, len(*userList))

	for _, user := range *userList {
		response = append(response, MapUserToUserResponse(&user))
	}
	return response
}
