package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	args := os.Args
	var env string

	if len(args) > 1 {
		env = strings.Split(args[1], "=")[1]
	} else {
		env = "local"
	}

	config.Setup(env)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Global.CorsUrl}

	router.Use(cors.New(corsConfig))

	router.GET("/health", handlers.HealthHandler)
	router.Run("localhost:" + strconv.Itoa(config.Global.Port))
}
