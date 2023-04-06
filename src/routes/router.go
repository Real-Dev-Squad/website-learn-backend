package routes

import "github.com/gin-gonic/gin"

func SetupRouter(env string, version string) *gin.Engine {
	router := gin.Default()
	UseRoutes(router, version)
	return router
}
