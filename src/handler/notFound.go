package handler

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"statusCode": fiber.StatusNotFound,
		"error":      "Not Found",
		"message":    "Not Found",
	}) // => 404 "Not Found"
}
