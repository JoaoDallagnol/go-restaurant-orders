package service

import (
	"fmt"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
)

type AuthService interface {
	RegisterUser(userReq *model.RegisterUserRequest) string
	Login(loginReq *model.UserLoginRequest) string
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) RegisterUser(userReq *model.RegisterUserRequest) string {
	return fmt.Sprintf("Usuario %s registrado com sucesso!", userReq.Name)
}

func (s *authService) Login(loginReq *model.UserLoginRequest) string {
	return "Usuario Logado!"
}
