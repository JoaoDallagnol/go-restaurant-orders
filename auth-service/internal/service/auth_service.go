package service

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
)

type AuthService interface {
	RegisterUser(userReq *model.RegisterUserRequest) model.UserResponse
	Login(loginReq *model.UserLoginRequest) string
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) RegisterUser(userReq *model.RegisterUserRequest) model.UserResponse {
	user := mapper.MapCreateUserRequestToUser(userReq)

	createdUser, err := s.userRepository.CreateUser(&user)
	if err != nil {
		panic("Falied to create user: " + err.Error())
	}

	return mapper.MapUserToUserResponse(createdUser)
}

func (s *authService) Login(loginReq *model.UserLoginRequest) string {
	return "Usuario Logado!"
}
