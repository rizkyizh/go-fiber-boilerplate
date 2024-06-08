package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"uniqueIndex;size:100;not null"`
	Age   int    `gorm:"default:0"`
}
