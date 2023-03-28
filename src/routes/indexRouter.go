package routes

import (
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/handlers"

	"github.com/gin-gonic/gin"
)

func RouteHandler(router *gin.Engine) {
	router.GET("/health", handlers.HealthHandler)
}
