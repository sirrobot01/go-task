package routers

import (
	"github.com/labstack/echo/v4"
	"task/handlers"
)

func InitTaskRoute(r *echo.Group) {
	r.GET("/:id", handlers.GetTasks)
	r.GET("/", handlers.GetTasks)
	r.POST("/", handlers.CreateTask)
	r.PUT("/:id", handlers.UpdateTask)
	r.DELETE("/:id", handlers.DeleteTask)
}
