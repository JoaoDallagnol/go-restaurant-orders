package service

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
)

type UserService interface {
	GetAllUser() []model.UserResponse
	GetUserById(userId string) string
	UpdateUser(userId string, userReq *model.RegisterUserRequest) string
	DeleteUser(userId string) string
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) GetAllUser() []model.UserResponse {
	userList, err := s.userRepository.GetAllUsers()

	if err != nil {
		panic("Falied to retrieve user list: " + err.Error())
	}

	if len(userList) == 0 {
		return []model.UserResponse{}
	}

	return mapper.MapUserListToUserResponseList(&userList)
}

func (s *userService) GetUserById(userId string) string {
	return "Usuario by Id"
}

func (s *userService) UpdateUser(userId string, userReq *model.RegisterUserRequest) string {
	return "Usuario atualizado"
}

func (s *userService) DeleteUser(userId string) string {
	return "Usuario deletado"
}
