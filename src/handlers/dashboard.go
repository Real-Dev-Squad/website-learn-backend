package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	isUserAuthenticated, _ := c.Get("Authenticated")
	if isUserAuthenticated == true {
		userId, _ := c.Get("userId")
		respMessage := "You have successfully reached dashboard user :" + userId.(string)
		c.IndentedJSON(http.StatusOK, gin.H{"message": respMessage})
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
	}
}
