package service

import (
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
	user := mapCreateUserRequestToUser(userReq)

	createdUser, err := s.userRepository.CreateUser(&user)
	if err != nil {
		panic("Falied to create user: " + err.Error())
	}

	return mapUserToUserResponse(createdUser)
}

func (s *authService) Login(loginReq *model.UserLoginRequest) string {
	return "Usuario Logado!"
}

func mapCreateUserRequestToUser(dto *model.RegisterUserRequest) model.User {
	return model.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func mapUserToUserResponse(user *model.User) model.UserResponse {
	return model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}
}
