package middleware

import (
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Global.CorsUrl,
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
	app.Use(helmet.New())
	// TODO: activate this later after implementation of authentication
	// app.Use(csrf.New())
	app.Use(logger.New())
}
