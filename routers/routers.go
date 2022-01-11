package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"task/common"
)

func InitRouters(e *echo.Echo) {

	TaskRoute(e.Group("tasks"))
	UserRoute(e.Group("/users"))

	admin := e.Group("/admin")
	admin.Use(middleware.JWTWithConfig(common.JWTConfig))
	AdminRoute(admin)

}
