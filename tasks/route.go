package tasks

import "github.com/labstack/echo/v4"

func InitRoute(r *echo.Group) {
	r.GET("/:id", getTask)
	r.GET("/", getTasks)
	r.POST("/", createTask)
	r.PUT("/:id", updateTask)
	r.DELETE("/:id", deleteTask)
}
