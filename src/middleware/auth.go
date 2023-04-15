package middleware

import (
	"fmt"
	"log"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/utils"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func verifyUser(context *gin.Context) (bool, bool, error) {

	authenticatedUser := false
	internalServerError := false

	cookie, err := context.Cookie(config.Global.CookieName)
	if err != nil {
		return authenticatedUser, internalServerError, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.Global.PublicKey))
	if err != nil {
		internalServerError = true
		return authenticatedUser, internalServerError, err
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) { return key, nil })
	if err != nil {
		return authenticatedUser, internalServerError, err
	}
	if !token.Valid {
		return authenticatedUser, internalServerError, fmt.Errorf("Invalid Token")
	}

	parsedJWT, _, err := new(jwt.Parser).ParseUnverified(cookie, jwt.MapClaims{})
	if err != nil {
		return authenticatedUser, internalServerError, err
	}

	claims, ok := parsedJWT.Claims.(jwt.MapClaims)

	if !ok {
		return authenticatedUser, internalServerError, err
	}

	context.Set("userId", claims["userId"])

	authenticatedUser = true
	return authenticatedUser, internalServerError, err
}

func Authenticate(context *gin.Context) {
	authenticatedUser, internalServerError, err := verifyUser(context)
	if err != nil {
		log.Println("Some Error occured")
	}
	if internalServerError {
		utils.InternalServerErrorResponse(context)
	} else if !authenticatedUser {
		utils.UnauthorizedResponse(context)
	} else {
		context.Next()
	}
}
