package common

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JwtClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

var JWTConfig = middleware.JWTConfig{
	Claims:     &JwtClaims{},
	SigningKey: []byte("secret"),
}
