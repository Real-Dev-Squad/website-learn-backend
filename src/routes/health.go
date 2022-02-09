package routes

import (
	"fmt"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func HealthRoute(app *fiber.App) {

	health := app.Group("health")

	health.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(config.Global.Port)
		return c.SendString("hello")
	})

	health.Get("/dashboard", monitor.New())

}
