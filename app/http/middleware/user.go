package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func GetUserId(c echo.Context) int64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(int64)
}
