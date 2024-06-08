package middlewares

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

func ValidateRequest(reqBody interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		typ := reflect.TypeOf(reqBody).Elem()
		v := reflect.New(typ).Interface()
		errHandler := &utils.ResponseHandler{}

		if err := c.BodyParser(v); err != nil {
			return errHandler.BadRequest(c, []string{"Cannot parse JSON"})
		}

		if err := utils.ValidateStruct(v); err != nil {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.Field()+": "+err.Tag())
			}
			return errHandler.BadRequest(c, errors)
		}

		c.Locals("validatedReqBody", v)
		return c.Next()
	}
}
