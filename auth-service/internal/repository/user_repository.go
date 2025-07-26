package repository

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserById(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) UpdateUser(user *model.User) (*model.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
