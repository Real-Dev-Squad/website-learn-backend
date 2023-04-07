package routes

import (
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/controllers"
	"github.com/gin-gonic/gin"
)

func UseRoutes(router *gin.Engine, version string) {
	ver := "/v" + version
	router.GET(ver+"/health", controllers.Health)
}
