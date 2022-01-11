package routers

import (
	"github.com/labstack/echo/v4"
	"task/handlers"
)

func TaskRoute(r *echo.Group) {
	r.GET("/:id", handlers.GetTask)
	r.GET("/", handlers.GetTasks)
	r.POST("/", handlers.CreateTask)
	r.PUT("/:id", handlers.UpdateTask)
	r.DELETE("/:id", handlers.DeleteTask)
}
