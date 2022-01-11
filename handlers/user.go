package handlers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"task/common"
	"task/models"
	"time"
)

func Register(e echo.Context) error {
	username := e.FormValue("username")
	password1 := e.FormValue("password1")
	password2 := e.FormValue("password2")

	if password1 != password2 {
		return e.JSON(http.StatusBadRequest, echo.Map{"message": "password1 and password2 must be the same"})
	}

	if user := models.Users[username]; user != nil {
		// User exists
		return e.JSON(http.StatusBadRequest, echo.Map{"message": "User with this username already exists"})
	}

	user := &models.User{
		Username: username,
		Password: password1,
		IsAdmin:  true,
	}
	models.Users[user.Username] = user
	return e.JSON(http.StatusOK, user)

}

func Login(e echo.Context) error {
	username := e.FormValue("username")
	password := e.FormValue("password")

	user := models.Users[username]

	if user == nil {
		return echo.ErrNotFound
	}

	if username != user.Username || password != user.Password {
		return echo.ErrUnauthorized
	}

	claims := &common.JwtClaims{
		Username: username,
		Admin:    true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, echo.Map{"token": t})
}
