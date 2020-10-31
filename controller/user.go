package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	models "github.com/hjcalderon10/bunny-backend/model"
	taskRepo "github.com/hjcalderon10/bunny-backend/repository/task"
	userRepo "github.com/hjcalderon10/bunny-backend/repository/user"
	taskServices "github.com/hjcalderon10/bunny-backend/service/task"
	userServices "github.com/hjcalderon10/bunny-backend/service/user"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	errors "github.com/hjcalderon10/bunny-backend/util/error"
	"github.com/hjcalderon10/bunny-backend/util/logger"
	"github.com/labstack/echo"
)

type User struct {
	BasePath     string
	UserPath     string
	UserTaskPath string
	userSrv      userServices.IService
	taskSrv      taskServices.ITaskService
}

func StartUser() User {
	resourcePath := fmt.Sprintf("/:%s", paramUserID)
	return User{
		BasePath:     "/users",
		UserPath:     resourcePath,
		UserTaskPath: fmt.Sprintf("%s/tasks", resourcePath),
		userSrv:      userServices.New(userRepo.New()),
		taskSrv:      taskServices.New(taskRepo.New()),
	}
}

func (ctrl User) GetAllUsers(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)

	users, err := ctrl.userSrv.GetAllUsers(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, users)
}

func (ctrl User) CreateUser(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := ctrl.userSrv.CreateUser(ctx, user); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.NoContent(http.StatusCreated)
}

func (ctrl User) ReadUser(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	idStr := c.Param(paramUserID)
	user := models.User{}

	id, err := strconv.Atoi(idStr)

	if err != nil || !(id > 0) {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, "ID must be a positive number"))
	}

	user.ID = uint16(id)

	if err := ctrl.userSrv.ReadUser(ctx, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, user)
}

func (ctrl User) UpdateUser(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	idStr := c.Param(paramUserID)
	user := models.User{}

	id, err := strconv.Atoi(idStr)

	if err != nil || !(id > 0) {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, "ID must be a positive number"))
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	user.ID = uint16(id)

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := ctrl.userSrv.UpdateUser(ctx, user); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.NoContent(http.StatusOK)
}

func (ctrl User) DeleteUser(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	idStr := c.Param(paramUserID)
	user := models.User{}

	id, err := strconv.Atoi(idStr)

	if err != nil || !(id > 0) {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, "ID must be a positive number"))
	}

	user.ID = uint16(id)

	if err := ctrl.userSrv.DeleteUser(ctx, user); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.NoContent(http.StatusOK)
}

func (ctrl User) CreateUserTask(c echo.Context) error {
	log := logger.New("-")
	ctx := context.WithValue(c.Request().Context(), settings.Commons.LogKey, log)
	idStr := c.Param(paramUserID)
	task := models.Task{}

	id, err := strconv.Atoi(idStr)

	if err != nil || !(id > 0) {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, "ID must be a positive number"))
	}

	task.UserID = uint16(id)

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, errors.HandleError(http.StatusBadRequest, err.Error()))
	}

	if err := ctrl.taskSrv.CreateTask(ctx, task); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.HandleError(http.StatusInternalServerError, err.Error()))
	}

	return c.NoContent(http.StatusCreated)
}
