package middleware

import (
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Global.Env,
		AllowCredentials: true,
	}))
	app.Use(csrf.New())
	app.Use(logger.New())
}
