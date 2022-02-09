package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/handler"
)

func SetupRoutes(app *fiber.App) {

	HealthRoute(app)

	app.Use(handler.NotFound)

}
