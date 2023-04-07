package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/routes"
)

func main() {
	value, keyExists := os.LookupEnv("env")
	port := strconv.Itoa(8000)
	version := strconv.Itoa(1)

	if keyExists {
		log.Printf("Running in %v mode\n", value)
	} else {
		log.Printf("Running in local dev mode\n")
		value = "local"
	}

	router := routes.SetupRouter(value, version)
	log.Printf("Server is running on http://localhost:%v\n", port)
	router.Run("localhost:" + port)
}
