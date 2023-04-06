package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "server is up and running"})
}
