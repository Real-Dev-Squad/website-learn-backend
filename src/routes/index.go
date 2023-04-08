package routes

import (
	"github.com/gin-gonic/gin"
)

func UseRoutes(router *gin.Engine, version string) {
	ver := "/v" + version
	indexGroup := router.Group(ver)
	{
		healthGroup := indexGroup.Group("/health")
		{
			HealthRoutes(healthGroup)
		}
	}
}
