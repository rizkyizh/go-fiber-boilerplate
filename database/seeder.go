package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/rizkyizh/go-fiber-boilerplate/app/models"
)

type seedUser struct {
	name     string
	email    string
	password string
	age      int
	role     string
}

var seedUsers = []seedUser{
	{name: "Admin", email: "admin@example.com", password: "Admin1234!", age: 30, role: models.RoleAdmin},
	{name: "User", email: "user@example.com", password: "User1234!", age: 25, role: models.RoleUser},
}

// Seed populates the database with initial fixture data.
func Seed() {
	for _, s := range seedUsers {
		var count int64
		DB.Model(&models.User{}).Where("email = ?", s.email).Count(&count)
		if count > 0 {
			continue
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(s.password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("seeder: failed to hash password for %s: %v", s.email, err)
			continue
		}

		user := &models.User{
			Name:     s.name,
			Email:    s.email,
			Password: string(hashed),
			Age:      s.age,
			Role:     s.role,
		}

		if err := DB.Create(user).Error; err != nil {
			log.Printf("seeder: failed to create user %s: %v", s.email, err)
			continue
		}

		log.Printf("seeder: created user %s (%s)", s.email, s.role)
	}
}
