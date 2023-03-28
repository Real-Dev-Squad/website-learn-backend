package main

import (
	"os"
	"strings"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/routes"
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

	routes.Setup()
}
