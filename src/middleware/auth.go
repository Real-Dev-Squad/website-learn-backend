package middleware

import (
	"fmt"
	"log"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/utils"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func verifyUser(context *gin.Context) (bool, error) {

	authenticatedUser := false

	cookie, err := context.Cookie(config.Global.CookieName)
	if err != nil {
		return authenticatedUser, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.Global.PublicKey))
	if err != nil {
		return authenticatedUser, err
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) { return key, nil })
	if err != nil {
		return authenticatedUser, err
	}
	if !token.Valid {
		return authenticatedUser, fmt.Errorf("Invalid Token")
	}

	parsedJWT, _, err := new(jwt.Parser).ParseUnverified(cookie, jwt.MapClaims{})
	if err != nil {
		return authenticatedUser, err
	}

	claims, ok := parsedJWT.Claims.(jwt.MapClaims)

	if !ok {
		return authenticatedUser, nil
	}

	context.Set("userId", claims["userId"])

	authenticatedUser = true
	return authenticatedUser, nil
}

func Authenticate(context *gin.Context) {
	authenticatedUser, err := verifyUser(context)
	if err != nil {
		log.Println("Some Error occured")
	}

	if !authenticatedUser {
		utils.UnauthorizedResponse(context)
	} else {
		context.Next()
	}
}
