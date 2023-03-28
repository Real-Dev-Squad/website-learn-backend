package routes

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
)

func Setup() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Global.CorsUrl}

	router.Use(cors.New(corsConfig))

	RouteHandler(router)

	router.Run("localhost:" + strconv.Itoa(config.Global.Port))
}
