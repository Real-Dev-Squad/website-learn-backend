package errorhandlerUtil

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Default error handler
var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.ErrInternalServerError.Code
	message := fiber.ErrInternalServerError.Message

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
		message = e.Message
	}
	// Return statuscode with error message
	log.Printf("Error occured: %v", err)
	return c.Status(code).JSON(fiber.Map{
		"status":  code,
		"error":   message,
		"message": err.Error(),
	})
}
