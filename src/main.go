package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/routes"
)

func main() {

	env, _ := os.LookupEnv("env")

	config.Setup(env)
	port := strconv.Itoa(config.Global.Port)
	version := strconv.Itoa(config.Global.Version)

	router := routes.SetupRouter(config.Global.Env, version)
	log.Printf("Server is running on http://localhost:%v/v%v/\n", port, version)
	router.Run("localhost:" + port)
}
