package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status string       `json:"status"`
	Error  ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Code    string   `json:"code"`
	Message []string `json:"message"`
}

type ResponseData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ResponseHandler struct{}

func (h *ResponseHandler) Ok(
	c *fiber.Ctx,
	data interface{},
	message string,
	meta *Meta,
) error {
	return c.Status(fiber.StatusOK).JSON(ResponseData{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func (h *ResponseHandler) Created(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusCreated).JSON(ResponseData{
		Status:  "Created",
		Message: message,
		Data:    data,
		Meta:    nil,
	})
}

func (h *ResponseHandler) BadRequest(c *fiber.Ctx, messages []string) error {
	return h.sendError(c, "BAD_REQUEST", fiber.StatusBadRequest, messages)
}

func (h *ResponseHandler) Forbidden(c *fiber.Ctx, messages []string) error {
	return h.sendError(c, "FORBIDDEN", fiber.StatusForbidden, messages)
}

func (h *ResponseHandler) NotFound(c *fiber.Ctx, messages []string) error {
	return h.sendError(c, "NOT_FOUND", fiber.StatusNotFound, messages)
}

func (h *ResponseHandler) InternalServerError(c *fiber.Ctx, messages []string) error {
	return h.sendError(c, "INTERNAL_SERVER_ERROR", fiber.StatusInternalServerError, messages)
}

func (h *ResponseHandler) sendError(
	c *fiber.Ctx,
	errorCodeMessage string,
	errorCode int,
	messages []string,
) error {
	response := ErrorResponse{
		Status: "error",
		Error: ErrorDetails{
			Code:    errorCodeMessage,
			Message: messages,
		},
	}

	return c.Status(errorCode).JSON(response)
}
