package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
)

func Setup() {

	if config.Global.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Global.CorsUrl}

	router.Use(cors.New(corsConfig))

	RouteHandler(router)

	fmt.Println("Server is running on Port:", config.Global.Port)
	router.Run("localhost:" + strconv.Itoa(config.Global.Port))
}
