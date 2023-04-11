package middleware

import (
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/utils"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	authenticatedUser := true

	cookie, err := context.Cookie(config.Global.CookieName)

	if err == nil {
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.Global.PublicKey))

		if err == nil {

			token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) { return key, nil })

			if err == nil && token.Valid {
				parsedJWT, _, err := new(jwt.Parser).ParseUnverified(cookie, jwt.MapClaims{})

				if err == nil {

					if claims, ok := parsedJWT.Claims.(jwt.MapClaims); ok {
						userId := claims["userId"]
						context.Set("userId", userId)
					}
					context.Next()
				} else {
					authenticatedUser = false
				}
			} else {
				authenticatedUser = false
			}
		} else {
			authenticatedUser = false
		}
	} else {
		authenticatedUser = false
	}

	if !authenticatedUser {
		utils.UnauthorizedResponse(context)
	}
}
