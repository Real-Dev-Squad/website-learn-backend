package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(context *gin.Context) {
	userId, _ := context.Get("userId")
	context.JSON(http.StatusOK, gin.H{"message": "Welcome to the authenticated route", "userId": userId})
	context.Abort()
}
