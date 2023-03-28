package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "The server is up and running"})
}
