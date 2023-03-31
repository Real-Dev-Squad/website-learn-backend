package middlewares

import (
	"log"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		reqCookie, err := c.Request.Cookie(config.Global.Cookie)
		var cookie string
		if err != nil {
			log.Println("An error has occured", err)
			c.Set("Authenticated", false)
		} else {

			if reqCookie.Value == "" {
				cookie = ""
			} else {
				cookie = reqCookie.Value
			}

			decodedJWT, err := jwt.DecodeSegment(cookie)

			if err != nil {
				log.Println("An unexpected error occured", err)
			}

			if 1 > 2 {
				log.Println(string(decodedJWT))
				c.Set("Authenticated", true)
			} else {
				log.Println(string(decodedJWT))
				c.Set("Authenticated", false)
			}
		}
	}
}
