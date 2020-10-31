package controller

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	models "github.com/hjcalderon10/bunny-backend/model"
	mocks "github.com/hjcalderon10/bunny-backend/test/mocks/service"
	testUtils "github.com/hjcalderon10/bunny-backend/test/util"
	"github.com/stretchr/testify/assert"
)

var (
	taskID    = "1"
	notTaskID = "notAnID"
	errTask   = fmt.Errorf("Ooops, this' awful")
)

func TestGetAllTasks(t *testing.T) {
	endPoint := "/api/tasks"

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("GetAllTasks").Return([]models.Task{}, nil)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.GetAllTasks(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "GetAllTasks", 1)
}

func TestGetAllTasksServiceError(t *testing.T) {
	endPoint := "/api/tasks"

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("GetAllTasks").Return(nil, errTask)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.GetAllTasks(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "GetAllTasks", 1)
}

func TestReadTask(t *testing.T) {
	endPoint := fmt.Sprintf("/api/tasks/%s", taskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	c.SetParamNames(paramTaskID)
	c.SetParamValues(taskID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("ReadTask").Return(nil)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.ReadTask(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "ReadTask", 1)
}

func TestReadTaskBindError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/tasks/%s", notTaskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	c.SetParamNames(paramTaskID)
	c.SetParamValues(notTaskID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("ReadTask").Return(nil)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.ReadTask(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "ReadTask", 0)
}

func TestReadTaskServiceError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/tasks/%s", taskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	c.SetParamNames(paramTaskID)
	c.SetParamValues(taskID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("ReadTask").Return(errTask)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.ReadTask(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "ReadTask", 1)
}

func TestUpdateTask(t *testing.T) {
	endPoint := fmt.Sprintf("/api/tasks/%s", taskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPut,
		endPoint,
		strings.NewReader(`{"title":"My task title"}`),
	)

	c.SetParamNames(paramTaskID)
	c.SetParamValues(taskID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("UpdateTask").Return(nil)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.UpdateTask(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "UpdateTask", 1)
}

func TestUpdateTaskBindError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/tasks/%s", notTaskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPut,
		endPoint,
		strings.NewReader(`{"title":"My task title"}`),
	)

	c.SetParamNames(paramTaskID)
	c.SetParamValues(notTaskID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("UpdateTask").Return(nil)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.UpdateTask(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "UpdateTask", 0)
}

func TestUpdateTaskServiceError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/tasks/%s", notTaskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPut,
		endPoint,
		strings.NewReader(`{"title":"My task title"}`),
	)

	c.SetParamNames(paramTaskID)
	c.SetParamValues(taskID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("UpdateTask").Return(errTask)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.UpdateTask(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "UpdateTask", 1)
}

func TestDeleteTask(t *testing.T) {
	endPoint := fmt.Sprintf("/api/tasks/%s", taskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodDelete,
		endPoint,
		nil,
	)

	c.SetParamNames(paramTaskID)
	c.SetParamValues(taskID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("DeleteTask").Return(nil)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.DeleteTask(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "DeleteTask", 1)
}

func TestDeleteTaskBindError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/tasks/%s", notTaskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodDelete,
		endPoint,
		nil,
	)

	c.SetParamNames(paramTaskID)
	c.SetParamValues(notTaskID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("DeleteTask").Return(nil)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.DeleteTask(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "DeleteTask", 0)
}

func TestDeleteTaskServiceError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/tasks/%s", taskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodDelete,
		endPoint,
		nil,
	)

	c.SetParamNames(paramTaskID)
	c.SetParamValues(taskID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("DeleteTask").Return(errTask)

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.DeleteTask(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "DeleteTask", 1)
}
