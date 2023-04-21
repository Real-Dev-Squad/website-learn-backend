package middleware

import (
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/utils"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type verifyResponse struct {
	authenticatedUser   bool
	internalServerError bool
}

func verifyUser(context *gin.Context) (response verifyResponse) {

	cookie, err := context.Cookie(config.Global.CookieName)
	if err != nil {
		return
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.Global.PublicKey))
	if err != nil {
		response.internalServerError = true
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) { return key, nil })
	if err != nil {
		return
	}
	if !token.Valid {
		return
	}

	parsedJWT, _, err := new(jwt.Parser).ParseUnverified(cookie, jwt.MapClaims{})
	if err != nil {
		return
	}

	claims, ok := parsedJWT.Claims.(jwt.MapClaims)

	if !ok {
		return
	}

	context.Set("userId", claims["userId"])

	response.authenticatedUser = true
	return
}

func Authenticate(context *gin.Context) {

	response := verifyUser(context)

	if response.internalServerError {
		utils.InternalServerErrorResponse(context)
	} else if !response.authenticatedUser {
		utils.UnauthorizedResponse(context)
	} else {
		context.Next()
	}
}
