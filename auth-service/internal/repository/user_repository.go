package repository

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/db"
	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
)

func CreateUser(user *model.User) error {
	return db.DB.Create(user).Error
}

func GetUserById(id uint) (*model.User, error) {
	var user model.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func DeleteUser(id uint) error {
	return db.DB.Delete(&model.User{}, id).Error
}
