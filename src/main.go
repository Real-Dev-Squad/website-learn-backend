package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/routes"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/utils"
)

func main() {

	env, _ := os.LookupEnv("env")

	config.Setup(env)
	port := strconv.Itoa(config.Global.Port)

	args := os.Args
	if len(args) > 1 {
		requiredArg := args[1]
		if requiredArg == "validate" {
			utils.ValidateSetup()
		}
	} else {
		router := routes.SetupRouter(config.Global.Env)
		log.Printf("Server is running on http://localhost:%v/\n", port)
		router.Run("localhost:" + port)
	}
}
