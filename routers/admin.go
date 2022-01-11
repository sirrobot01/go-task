package routers

import (
	"github.com/labstack/echo/v4"
	"task/handlers"
)

func AdminRoute(r *echo.Group) {
	r.GET("/logged", handlers.GetLoggedInUser)
}
