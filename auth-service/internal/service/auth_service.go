package service

import (
	"fmt"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
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
	hashedPaswword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		fmt.Println("Error hashing password:", hashErr.Error())
	}

	user.Password = string(hashedPaswword)

	createdUser, err := s.userRepository.CreateUser(&user)
	if err != nil {
		panic("Falied to create user: " + err.Error())
	}

	return mapper.MapUserToUserResponse(createdUser)
}

func (s *authService) Login(loginReq *model.UserLoginRequest) string {
	return "Usuario Logado!"
}
