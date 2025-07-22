package service

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
)

type UserService interface {
	GetAllUser() string
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

func (s *userService) GetAllUser() string {
	return "Lista de usuario"
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
