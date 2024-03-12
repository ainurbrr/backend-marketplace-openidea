package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
)

var IsLoggedIn = echojwt.WithConfig(echojwt.Config{
	SigningMethod: "HS256",
	SigningKey:    []byte("S3cret"),
})

func CreateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte("S3ret")
	return token.SignedString(byteSecret)
}
