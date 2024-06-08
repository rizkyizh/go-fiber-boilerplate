package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/app/controllers"
	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/repositories"
	"github.com/rizkyizh/go-fiber-boilerplate/app/services"
	"github.com/rizkyizh/go-fiber-boilerplate/middlewares"
)

func UserRoutes(route fiber.Router) {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	route.Get("/", userController.GetUsers)
	route.Get("/:id", userController.GetUser)
	route.Post("/", middlewares.ValidateRequest(&dto.UserDTO{}), userController.CreateUser)
	route.Patch(
		"/:id",
		middlewares.ValidateRequest(&dto.UpdateUserDTO{}),
		userController.UpdateUser,
	)
	route.Delete("/:id", userController.DeleteUser)
}
