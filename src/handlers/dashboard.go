package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	isUserAuthenticated, _ := c.Get("Authenticated")
	if isUserAuthenticated == true {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "you reached the dashboard"})
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
	}
}
