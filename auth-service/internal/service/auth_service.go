package service

import (
	"errors"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	RegisterUser(userReq *model.RegisterUserRequest) (model.UserResponse, error)
	Login(loginReq *model.UserLoginRequest) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) RegisterUser(userReq *model.RegisterUserRequest) (model.UserResponse, error) {
	user := mapper.MapCreateUserRequestToUser(userReq)
	hashedPaswword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.UserResponse{}, errs.NewInternalError(err.Error())
	}

	user.Password = string(hashedPaswword)

	createdUser, err := s.userRepository.CreateUser(&user)
	if err != nil {
		return model.UserResponse{}, errs.NewInternalError(err.Error())
	}

	return mapper.MapUserToUserResponse(createdUser), nil
}

func (s *authService) Login(loginReq *model.UserLoginRequest) (string, error) {

	user, err := s.userRepository.GetUserByEmail(loginReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "model.UserResponse{}", errs.NewAuthInvalidCredentials()
		}
		return "", errs.NewInternalError(err.Error())
	}

	return "%s, Usuario Logado!" + user.Email, nil
}
