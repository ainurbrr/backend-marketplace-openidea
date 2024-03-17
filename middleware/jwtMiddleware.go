package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)
var secret = os.Getenv("JWT_SECRET")

var IsLoggedIn = echojwt.WithConfig(echojwt.Config{
	SigningMethod: "HS256",
	SigningKey:    []byte(secret),
})

func CreateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(secret)
	return token.SignedString(byteSecret)
}

func ExtractUserID(c echo.Context) (userId string, err error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return userId, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	userId = claims["userId"].(string)

	return userId, nil
}
