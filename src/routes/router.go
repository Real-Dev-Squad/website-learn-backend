package routes

import "github.com/gin-gonic/gin"

func SetupRouter(env string) *gin.Engine {
	router := gin.Default()
	UseRoutes(router)
	return router
}
