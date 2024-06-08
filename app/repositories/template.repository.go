package repositories

import (
	"github.com/rizkyizh/go-fiber-boilerplate/app/models"
	"github.com/rizkyizh/go-fiber-boilerplate/database"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUsers(page, perPage int) ([]models.User, int64, error)
	GetUser(userId uint) (*models.User, error)
	UpdateUser(userId uint, user *models.User) error
	DeleteUser(userId uint) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user *models.User) error {
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUsers(page, perPage int) ([]models.User, int64, error) {
	var users []models.User
	var totalItems int64

	// count total items
	database.DB.Model(&models.User{}).Count(&totalItems)

	if err := database.DB.Offset((page - 1) * perPage).Limit(perPage).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, totalItems, nil
}

func (r *userRepository) GetUser(userId uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(userId uint, user *models.User) error {
	if err := database.DB.Model(&models.User{}).Where("id = ?", userId).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(id uint) error {
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
