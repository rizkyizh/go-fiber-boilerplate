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

// Get List of Users godoc
// @Summary Get List of Users
// @Description Get List of Users
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param perPage query int false "PerPage"
// @Param sort query string false "Sort"
// @Param search query string false "Search"
// @Param status query string false "Status"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /users [get]
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

// Create User godoc
// @Summary Create User
// @Description Create User
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDTO true "User"
// @Success 201 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /users [post]
func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	var dto dto.CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	err := ctrl.service.CreateUser(&dto)
	if err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}
	return h.Created(c, nil, "user created successfully")
}

// Get User godoc
// @Summary Get User
// @Description Get User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/{id} [get]
func (ctrl *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h := &utils.ResponseHandler{}

	user, err := ctrl.service.GetUserById(id)
	if err != nil {
		return h.NotFound(c, []string{err.Error()})
	}

	return h.Ok(c, user, "users fetched successfully", nil)
}

// Update User godoc
// @Summary Update User
// @Description Update User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param user body dto.UpdateUserDTO true "User"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h := &utils.ResponseHandler{}

	var dto *dto.UpdateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	user, err := ctrl.service.UpdateUser(id, dto)
	if err != nil {
		return h.InternalServerError(c, []string{err.Error()})
	}
	return h.Ok(c, user, "user updated successfully", nil)
}

// Delete User godoc
// @Summary Delete User
// @Description Delete User
// @Tags Users
// @Produce json
// @Accept json
// @Param id path string true "ID"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/{id} [delete]
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
