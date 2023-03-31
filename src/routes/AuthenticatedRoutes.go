package routes

import (
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/handlers"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/middlewares"
	"github.com/gin-gonic/gin"
)

func UseAuthenticatedRoutes(router *gin.Engine) {
	router.Use(middlewares.Auth())

	router.GET("/health/dashboard", handlers.Dashboard)
}
