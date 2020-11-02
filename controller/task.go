package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	models "github.com/hjcalderon10/bunny-backend/model"
	taskRepo "github.com/hjcalderon10/bunny-backend/repository/task"
	services "github.com/hjcalderon10/bunny-backend/service/task"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	errors "github.com/hjcalderon10/bunny-backend/util/error"
	"github.com/hjcalderon10/bunny-backend/util/logger"
	"github.com/labstack/echo"
)

type Task struct {
	BasePath   string
	TaskPath   string
	StatesPath string
	service    services.IService
}

func StartTask() Task {
	return Task{
		BasePath:   "/tasks",
		TaskPath:   fmt.Sprintf("/:%s", paramTaskID),
		StatesPath: "/states",
		service:    services.New(taskRepo.New()),
	}
}

func (ctrl Task) GetAllTasks(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)

	tasks, err := ctrl.service.GetAllTasks(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, tasks)
}

func (ctrl Task) GetAllTaskStates(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)

	taskStates, err := ctrl.service.GetAllTaskStates(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, taskStates)
}

func (ctrl Task) ReadTask(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	idStr := c.Param(paramTaskID)
	task := models.Task{}

	id, err := strconv.Atoi(idStr)

	if err != nil || !(id > 0) {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, "ID must be a positive number"))
	}

	task.ID = uint16(id)

	if err := ctrl.service.ReadTask(ctx, &task); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, task)

}

func (ctrl Task) UpdateTask(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	idStr := c.Param(paramTaskID)
	task := models.Task{}

	id, err := strconv.Atoi(idStr)

	if err != nil || !(id > 0) {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, "ID must be a positive number"))
	}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	task.ID = uint16(id)

	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := ctrl.service.UpdateTask(ctx, task); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.NoContent(http.StatusOK)
}

func (ctrl Task) DeleteTask(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	idStr := c.Param(paramTaskID)
	task := models.Task{}

	id, err := strconv.Atoi(idStr)

	if err != nil || !(id > 0) {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, "ID must be a positive number"))
	}

	task.ID = uint16(id)

	if err := ctrl.service.DeleteTask(ctx, task); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.NoContent(http.StatusOK)
}
