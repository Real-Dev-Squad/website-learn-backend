package routes

import (
	"github.com/gofiber/fiber/v2"

	resourceHandler "github.com/Real-Dev-Squad/gopher-cloud-service/src/handler/resource"
	validator "github.com/Real-Dev-Squad/gopher-cloud-service/src/utils/validator"
)

func ResourcesRoute(app *fiber.App) {
	// TODO: discuss about the route names and grouping for routes
	resources := app.Group("learn")

	resources.Get("/resources", resourceHandler.GetResources)

	resources.Get("/resources/:id", resourceHandler.GetResource)

	resources.Post("/resources", validator.PostResource, resourceHandler.PostResource)

}
