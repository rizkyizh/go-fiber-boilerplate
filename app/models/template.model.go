package models

import "gorm.io/gorm"

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100;not null"`
	Email        string `gorm:"uniqueIndex;size:100;not null"`
	Age          int    `gorm:"default:0"`
	Password     string `gorm:"size:255;not null"`
	Role         string `gorm:"size:20;not null;default:'user'"`
	RefreshToken string `gorm:"size:512"`
}
