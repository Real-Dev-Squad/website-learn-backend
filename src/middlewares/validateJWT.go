package middlewares

import (
	"log"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func ValidateJWT(cookie string) bool {
	println("Inside validate JWT function")
	println("Received the following cookie", cookie)
	os.Setenv("PUBLIC_KEY", "<PUBLIC_KEY_TO_BE_SET>")
	// available claims = userId, iat, exp
	key, available := os.LookupEnv("PUBLIC_KEY")
	if available {
		parts := strings.Split(cookie, ".")
		err := jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], []byte(key))
		if err != nil {
			log.Print("unable to verify", err)
			return false
		} else {
			return true
		}
	}
	return false
}

// func keyFunc() {
// 	return "hello"
// }
