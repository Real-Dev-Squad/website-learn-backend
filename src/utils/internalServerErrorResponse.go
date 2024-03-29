package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerErrorResponse(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{"message": "An internal server error occurred"})
	context.Abort()
}
