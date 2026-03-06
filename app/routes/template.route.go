package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/app/controllers"
	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/models"
	"github.com/rizkyizh/go-fiber-boilerplate/app/repositories"
	"github.com/rizkyizh/go-fiber-boilerplate/app/services"
	"github.com/rizkyizh/go-fiber-boilerplate/middlewares"
)

func UserRoutes(route fiber.Router) {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	route.Get("/", middlewares.AuthRequired(), userController.GetUsers)
	route.Get("/:id", middlewares.AuthRequired(), userController.GetUser)
	route.Post(
		"/",
		middlewares.AuthRequired(),
		middlewares.RequireRole(models.RoleAdmin),
		middlewares.ValidateRequest(&dto.CreateUserDTO{}),
		userController.CreateUser,
	)
	route.Patch(
		"/:id",
		middlewares.AuthRequired(),
		middlewares.RequireRole(models.RoleAdmin),
		middlewares.ValidateRequest(&dto.UpdateUserDTO{}),
		userController.UpdateUser,
	)
	route.Delete(
		"/:id",
		middlewares.AuthRequired(),
		middlewares.RequireRole(models.RoleAdmin),
		userController.DeleteUser,
	)
}
