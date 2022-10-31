package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/middleware"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/routes"
	errorHandler "github.com/Real-Dev-Squad/gopher-cloud-service/src/utils/errorHandler"
)

var Config string

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler.DefaultErrorHandler,
	})

	middleware.SetupMiddleware(app)

	routes.SetupRoutes(app)
	port := ":" + config.Global.Port
	log.Fatal(app.Listen(port))
}
