package service

import (
	"fmt"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) RegisterUser(userReq model.RegisterUserRequest) string {
	return fmt.Sprintf("Usuario %s registrado com sucesso!", userReq.Name)
}

func (s *AuthService) Login(loginReq model.UserLoginRequest) string {
	return "Usuario Logado!"
}
