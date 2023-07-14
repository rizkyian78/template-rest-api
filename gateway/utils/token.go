package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(payloadClaim *JWTClaims) (string, int, error) {
	seconds := 3600
	duration := time.Duration(seconds) * time.Second
	secret := "shh"

	payload := JWTClaims{
		Username: "rizkyian78",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(duration)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Println(err, "<<<<")
		// If there is an error in creating the JWT return an internal server error
		return "", 0, err
	}

	return tokenString, seconds, nil
}
