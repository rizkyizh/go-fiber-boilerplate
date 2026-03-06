package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/services"
	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new account with name, email, password, and age
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.RegisterDTO true "Register payload"
// @Success 201 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/register [post]
func (ctrl *AuthController) Register(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	payload := c.Locals("validatedReqBody").(*dto.RegisterDTO)

	if err := ctrl.service.Register(payload); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}
	return h.Created(c, nil, "registered successfully")
}

// Login godoc
// @Summary Login
// @Description Authenticate with email and password, returns access + refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.LoginDTO true "Login payload"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/login [post]
func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	payload := c.Locals("validatedReqBody").(*dto.LoginDTO)

	tokens, err := ctrl.service.Login(payload)
	if err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}
	return h.Ok(c, tokens, "login successful", nil)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Use a valid refresh token to obtain a new token pair
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.RefreshTokenDTO true "Refresh token payload"
// @Success 200 {object} utils.ResponseData
// @Failure 401 {object} utils.ErrorResponse
// @Router /auth/refresh [post]
func (ctrl *AuthController) RefreshToken(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	payload := c.Locals("validatedReqBody").(*dto.RefreshTokenDTO)

	tokens, err := ctrl.service.RefreshToken(payload.RefreshToken)
	if err != nil {
		return h.Unauthorized(c, []string{err.Error()})
	}
	return h.Ok(c, tokens, "token refreshed", nil)
}

// Logout godoc
// @Summary Logout
// @Description Invalidate the current user's refresh token
// @Tags Auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.ResponseData
// @Failure 401 {object} utils.ErrorResponse
// @Router /auth/logout [post]
func (ctrl *AuthController) Logout(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	claims := c.Locals("claims").(*utils.TokenClaims)

	if err := ctrl.service.Logout(claims.UserID); err != nil {
		return h.InternalServerError(c, []string{err.Error()})
	}
	return h.Ok(c, nil, "logged out successfully", nil)
}
