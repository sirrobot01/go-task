package handlers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"task/common"
)

func GetLoggedInUser(e echo.Context) error {
	user := e.Get("user")
	println(user)

	claims := user.(*jwt.Token).Claims.(*common.JwtClaims)

	println(claims)

	username := claims.Username

	return e.String(http.StatusOK, "Welcome "+username+"!")
}
