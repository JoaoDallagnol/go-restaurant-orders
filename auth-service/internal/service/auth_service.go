package service

import (
	"errors"
	"time"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	RegisterUser(userReq *model.RegisterUserRequest) (model.UserResponse, error)
	Login(loginReq *model.UserLoginRequest) (model.UserLoginResponse, error)
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

func (s *authService) Login(loginReq *model.UserLoginRequest) (model.UserLoginResponse, error) {
	user, err := s.userRepository.GetUserByEmail(loginReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.UserLoginResponse{}, errs.NewAuthInvalidCredentials()
		}
		return model.UserLoginResponse{}, errs.NewInternalError(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return model.UserLoginResponse{}, errs.NewAuthInvalidCredentials()
	}

	expiration := time.Duration(config.AppConfig.Auth.ExpirationMinutes) * time.Minute
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(expiration).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(config.AppConfig.Auth.Secret)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return model.UserLoginResponse{}, errs.NewInternalError(err.Error())
	}

	return buildLoginResponse(signedToken, expiration.String()), nil
}

func buildLoginResponse(token, expiration string) model.UserLoginResponse {
	return model.UserLoginResponse{
		Token:     token,
		ExpiresIn: expiration,
	}
}
