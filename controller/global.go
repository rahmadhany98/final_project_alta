package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func ReturnID(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	auth := claims["userId"].(float64)
	id := int(auth)
	return id
}
