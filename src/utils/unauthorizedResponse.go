package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnauthorizedResponse(context *gin.Context) {
	context.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthenticated user"})
	context.Abort()
}
