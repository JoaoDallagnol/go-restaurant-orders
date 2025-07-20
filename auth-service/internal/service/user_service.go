package service

import "github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUser() string {
	return "Lista de usuario"
}

func (s *UserService) GetUserById(userId string) string {
	return "Usuario by Id"
}

func (s *UserService) UpdateUser(userId string, userReq model.RegisterUserRequest) string {
	return "Usuario atualizado"
}

func (s *UserService) DeleteUser(userId string) string {
	return "Usuario deletado"
}
