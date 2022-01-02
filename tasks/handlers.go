package tasks

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"task/utils"
)

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, Tasks)
}

func getTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if task := Tasks[id]; task != nil {
		return c.JSON(http.StatusOK, task)
	}
	return c.JSON(http.StatusNotFound, utils.NotFound())

}

func updateTask(c echo.Context) error {
	task := new(Task)

	if err := c.Bind(task); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if task := Tasks[id]; task != nil {
		Tasks[id].Name = task.Name
		Tasks[id].Day = task.Day
		return c.JSON(http.StatusOK, Tasks[id])
	}
	return c.JSON(http.StatusNotFound, utils.NotFound())

}

func createTask(c echo.Context) error {
	task := &Task{
		ID: Seq,
	}
	if err := c.Bind(task); err != nil {
		return err
	}
	Tasks[task.ID] = task
	Seq++
	return c.JSON(http.StatusCreated, task)
}

func deleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if task := Tasks[id]; task != nil {
		delete(Tasks, id)
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusNotFound, utils.NotFound())
}
