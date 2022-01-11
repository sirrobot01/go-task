package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"task/common"
	"task/models"
)

func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Tasks)
}

func GetTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if task := models.Tasks[id]; task != nil {
		return c.JSON(http.StatusOK, task)
	}
	return c.JSON(http.StatusNotFound, common.NotFound)

}

func UpdateTask(c echo.Context) error {
	task := new(models.Task)

	if err := c.Bind(task); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if task := models.Tasks[id]; task != nil {
		models.Tasks[id].Name = task.Name
		models.Tasks[id].Day = task.Day
		return c.JSON(http.StatusOK, models.Tasks[id])
	}
	return c.JSON(http.StatusNotFound, common.NotFound)

}

func CreateTask(c echo.Context) error {
	task := &models.Task{
		ID: models.TaskSeq,
	}
	if err := c.Bind(task); err != nil {
		return err
	}
	models.Tasks[task.ID] = task
	models.TaskSeq++
	return c.JSON(http.StatusCreated, task)
}

func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if task := models.Tasks[id]; task != nil {
		delete(models.Tasks, id)
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusNotFound, common.NotFound)
}
