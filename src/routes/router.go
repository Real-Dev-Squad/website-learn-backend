package routes

import "github.com/gin-gonic/gin"

func SetupRouter(env string) *gin.Engine {
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	UseRoutes(router)
	return router
}
