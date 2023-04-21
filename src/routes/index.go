package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Real-Dev-Squad/website-learn-backend/src/config"
)

func UseRoutes(router *gin.Engine) {
	ver := "/v" + strconv.Itoa(config.Global.Version)
	indexGroup := router.Group(ver)
	healthGroup := indexGroup.Group("/health")
	HealthRoutes(healthGroup)
}
