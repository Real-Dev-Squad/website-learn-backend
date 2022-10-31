package validator

import (
	resourceSchema "github.com/Real-Dev-Squad/gopher-cloud-service/src/schema/resource"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func PostResource(c *fiber.Ctx) error {
	var errors []*ErrorResponse
	body := new(resourceSchema.Resources)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el ErrorResponse
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
