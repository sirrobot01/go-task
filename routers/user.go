package routers

import (
	"github.com/labstack/echo/v4"
	"task/handlers"
)

func UserRoute(r *echo.Group) {
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
}
