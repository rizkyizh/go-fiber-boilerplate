package repositories

import (
	"github.com/rizkyizh/go-fiber-boilerplate/app/models"
	"github.com/rizkyizh/go-fiber-boilerplate/database"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	UpdateRefreshToken(userID uint, refreshToken string) error
	ClearRefreshToken(userID uint) error
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) UpdateRefreshToken(userID uint, refreshToken string) error {
	return database.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("refresh_token", refreshToken).Error
}

func (r *authRepository) ClearRefreshToken(userID uint) error {
	return database.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("refresh_token", "").Error
}
