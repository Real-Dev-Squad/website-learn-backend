package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func HealthRoute(app *fiber.App) {

	health := app.Group("health")

	health.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	health.Get("/dashboard", monitor.New())

}
