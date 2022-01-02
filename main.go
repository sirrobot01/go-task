package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"task/tasks"
)

func main() {
	e := echo.New()

	// Middlewares

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Root

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Working")
	})

	tasks.InitRoute(e.Group("tasks"))

	e.Logger.Fatal(e.Start(":8080"))
}
