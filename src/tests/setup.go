package tests

import (
	"strconv"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/routes"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine
var Ver string
var Version string

func TestSetup() {
	config.Setup("test")
	Version = strconv.Itoa(config.Global.Version)
	Ver = "/v" + Version
	Router = routes.SetupRouter(config.Global.Env, Version)
}
