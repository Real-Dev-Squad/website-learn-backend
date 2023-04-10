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

	router := routes.SetupRouter(config.Global.Env)
	log.Printf("Server is running on http://localhost:%v/\n", port)
	router.Run("localhost:" + port)
}
