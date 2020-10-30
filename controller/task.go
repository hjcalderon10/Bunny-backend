package controller

import (
	"context"
	"net/http"

	models "github.com/hjcalderon10/bunny-backend/model"
	services "github.com/hjcalderon10/bunny-backend/service/task"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	errors "github.com/hjcalderon10/bunny-backend/util/error"
	"github.com/hjcalderon10/bunny-backend/util/logger"
	"github.com/labstack/echo"
)

type Task struct {
	BasePath string
	TaskPath string
	service  services.IService
}

func StartTask() Task {
	return Task{
		BasePath: "/tasks",
		TaskPath: "/:id",
		service:  services.New(),
	}
}

func (ctrl Task) ReadTask(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	task := models.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := ctrl.service.ReadTask(ctx, &task); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, task)
}

func (ctrl Task) UpdateTask(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	task := models.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := ctrl.service.UpdateTask(ctx, &task); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, task)
}

func (ctrl Task) DeleteTask(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	task := models.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := ctrl.service.DeleteTask(ctx, task); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.NoContent(http.StatusOK)
}
