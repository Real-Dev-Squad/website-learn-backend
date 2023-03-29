package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "you reached the dashboard"})
}
