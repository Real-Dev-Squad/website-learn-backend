package routes

import (
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/controllers"
	"github.com/gin-gonic/gin"
)

func HealthRoutes(router *gin.RouterGroup) {
	router.GET("", controllers.Health)
}
