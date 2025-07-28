package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/errs"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/mapper"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUser() []model.UserResponse
	GetUserById(id string) (model.UserResponse, error)
	UpdateUser(id string, userReq *model.RegisterUserRequest) model.UserResponse
	DeleteUser(id string)
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

func (s *userService) GetUserById(id string) (model.UserResponse, error) {
	userId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic("Invalid Id: " + err.Error())
	}

	user, err := s.userRepository.GetUserById(uint(userId))

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.UserResponse{}, &errs.UserNotFoundError{UserId: uint(userId)}
		}
		return model.UserResponse{}, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return mapper.MapUserToUserResponse(user), nil
}

func (s *userService) UpdateUser(id string, userReq *model.RegisterUserRequest) model.UserResponse {
	userId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic("Invalid ID: " + err.Error())
	}

	existingUser, err := s.userRepository.GetUserById(uint(userId))
	if err != nil {
		panic("User not found: " + err.Error())
	}

	hashedPaswword, hashErr := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		fmt.Println("Error hashing password:", hashErr.Error())
	}

	existingUser.Name = userReq.Name
	existingUser.Email = userReq.Email
	existingUser.Password = string(hashedPaswword)

	updatedUser, err := s.userRepository.UpdateUser(existingUser)
	if err != nil {
		panic("Failed to update user: " + err.Error())
	}

	return mapper.MapUserToUserResponse(updatedUser)
}

func (s *userService) DeleteUser(id string) {
	userId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic("Invalid Id: " + err.Error())
	}

	userFound, errUser := s.userRepository.GetUserById(uint(userId))

	if errUser != nil {
		panic("Failed to retrieve user: " + errUser.Error())
	}

	deleteErr := s.userRepository.DeleteUser(userFound)

	if deleteErr != nil {
		panic("Error on delete user: " + deleteErr.Error())
	}
}
