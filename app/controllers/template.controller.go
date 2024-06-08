package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/services"
	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (ctrl *UserController) GetUsers(c *fiber.Ctx) error {
	q := new(utils.QueryParams)
	if err := c.QueryParser(q); err != nil {
		return err
	}

	h := &utils.ResponseHandler{}

	users, meta, err := ctrl.service.GetAllUsers(*q)
	if err != nil {
		return h.InternalServerError(c, []string{err.Error()})
	}
	return h.Ok(c, users, "users fetched successfully", &meta)
}

func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	var dto dto.UserDTO
	if err := c.BodyParser(&dto); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	user, err := ctrl.service.CreateUser(dto)
	if err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}
	return h.Created(c, user, "user created successfully")
}

func (ctrl *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h := &utils.ResponseHandler{}

	user, err := ctrl.service.GetUserById(id)
	if err != nil {
		return h.NotFound(c, []string{err.Error()})
	}

	return h.Ok(c, user, "users fetched successfully", nil)
}

func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h := &utils.ResponseHandler{}

	var dto dto.UpdateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	user, err := ctrl.service.UpdateUser(id, dto)
	if err != nil {
		return h.InternalServerError(c, []string{err.Error()})
	}
	return h.Ok(c, user, "user updated successfully", nil)
}

func (ctrl *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h := &utils.ResponseHandler{}
	err := ctrl.service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return h.NotFound(c, []string{err.Error()})
		}
		return h.InternalServerError(c, []string{err.Error()})
	}

	return h.Ok(c, nil, "User deleted successfully", nil)
}
