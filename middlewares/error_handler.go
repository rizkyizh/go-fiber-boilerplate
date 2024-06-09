package middlewares

import (
	"encoding/json"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

func ValidateRequest(reqBody interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		errHandler := &utils.ResponseHandler{}

		// Decode request body into a map
		var bodyMap map[string]interface{}
		if err := json.Unmarshal(c.Body(), &bodyMap); err != nil {
			return errHandler.BadRequest(c, []string{"Invalid request body"})
		}

		typ := reflect.TypeOf(reqBody).Elem()
		v := reflect.New(typ).Interface()

		if err := c.BodyParser(v); err != nil {
			return errHandler.BadRequest(c, []string{"Cannot parse JSON"})
		}

		// Create a set of allowed fields from the DTO
		allowedFields := make(map[string]struct{})
		value := reflect.ValueOf(reqBody).Elem()
		for i := 0; i < value.NumField(); i++ {
			allowedFields[value.Type().Field(i).Tag.Get("json")] = struct{}{}
		}

		// Check for unexpected fields
		for field := range bodyMap {
			if _, found := allowedFields[field]; !found {
				return errHandler.BadRequest(c, []string{"Unexpected field: " + field})
			}
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
