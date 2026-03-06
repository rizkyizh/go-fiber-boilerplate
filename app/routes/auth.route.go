package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/app/controllers"
	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/repositories"
	"github.com/rizkyizh/go-fiber-boilerplate/app/services"
	"github.com/rizkyizh/go-fiber-boilerplate/middlewares"
)

func AuthRoutes(route fiber.Router) {
	authRepository := repositories.NewAuthRepository()
	userRepository := repositories.NewUserRepository()
	authService := services.NewAuthService(authRepository, userRepository)
	authController := controllers.NewAuthController(authService)

	route.Post("/register", middlewares.ValidateRequest(&dto.RegisterDTO{}), authController.Register)
	route.Post("/login", middlewares.ValidateRequest(&dto.LoginDTO{}), authController.Login)
	route.Post("/refresh", middlewares.ValidateRequest(&dto.RefreshTokenDTO{}), authController.RefreshToken)
	route.Post("/logout", middlewares.AuthRequired(), authController.Logout)
}
