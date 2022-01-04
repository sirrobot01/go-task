package routers

import "github.com/labstack/echo/v4"

func InitRouters(e *echo.Echo) {

	TaskRoute(e.Group("tasks"))

}
