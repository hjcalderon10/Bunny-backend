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
	userID    = "1"
	notUserID = "notAnID"
	errUser   = fmt.Errorf("Ooops, this' awful")
)

func TestGetAllUsers(t *testing.T) {
	endPoint := "/api/users"

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("GetAllUsers").Return([]models.User{}, nil)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.GetAllUsers(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "GetAllUsers", 1)
}

func TestGetAllUsersServiceError(t *testing.T) {
	endPoint := "/api/users"

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("GetAllUsers").Return(nil, errUser)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.GetAllUsers(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "GetAllUsers", 1)
}

func TestCreateUser(t *testing.T) {
	endPoint := "/api/users"

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPost,
		endPoint,
		strings.NewReader(`{"name":"Juanito Test"}`),
	)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("CreateUser").Return(models.User{}, nil)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.CreateUser(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "CreateUser", 1)
}

func TestCreateUserBindError(t *testing.T) {
	endPoint := "/api/users"

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPost,
		endPoint,
		strings.NewReader(`{"id":"Juanito Test"}`),
	)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("CreateUser").Return(models.User{}, nil)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.CreateUser(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "CreateUser", 0)
}

func TestCreateUserServiceError(t *testing.T) {
	endPoint := "/api/users"

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPost,
		endPoint,
		strings.NewReader(`{"name":"Juanito Test"}`),
	)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("CreateUser").Return(models.User{}, errUser)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.CreateUser(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "CreateUser", 1)
}

func TestReadUser(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s", userID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(userID)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("ReadUser").Return(nil)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.ReadUser(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "ReadUser", 1)
}

func TestReadUserBindError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s", notUserID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(notUserID)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("ReadUser").Return(nil)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.ReadUser(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "ReadUser", 0)
}

func TestReadUserServiceError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s", userID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(userID)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("ReadUser").Return(errUser)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.ReadUser(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "ReadUser", 1)
}

func TestUpdateUser(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s", userID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPut,
		endPoint,
		strings.NewReader(`{"name":"Juanito Test"}`),
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(userID)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("UpdateUser").Return(nil)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.UpdateUser(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "UpdateUser", 1)
}

func TestUpdateUserBindError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s", notUserID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPut,
		endPoint,
		strings.NewReader(`{"name":"Juanito Test"}`),
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(notUserID)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("UpdateUser").Return(nil)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.UpdateUser(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "UpdateUser", 0)
}

func TestUpdateUserServiceError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s", notUserID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPut,
		endPoint,
		strings.NewReader(`{"name":"Juanito Test"}`),
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(userID)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("UpdateUser").Return(errUser)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.UpdateUser(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "UpdateUser", 1)
}

func TestDeleteUser(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s", userID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodDelete,
		endPoint,
		nil,
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(userID)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("DeleteUser").Return(nil)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.DeleteUser(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "DeleteUser", 1)
}

func TestDeleteUserBindError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s", notUserID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodDelete,
		endPoint,
		nil,
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(notUserID)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("DeleteUser").Return(nil)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.DeleteUser(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "DeleteUser", 0)
}

func TestDeleteUserServiceError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s", userID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodDelete,
		endPoint,
		nil,
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(userID)

	userSrvMock := mocks.UserMock{}
	userSrvMock.On("DeleteUser").Return(errUser)

	ctrl := StartUser()
	ctrl.userSrv = &userSrvMock
	ctrl.DeleteUser(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	userSrvMock.AssertNumberOfCalls(t, "DeleteUser", 1)
}

func TestCreateUserTask(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s/tasks", userID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPost,
		endPoint,
		strings.NewReader(`{"title":"Juanito's first task"}`),
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(userID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("CreateTask").Return(models.Task{}, nil)

	ctrl := StartUser()
	ctrl.taskSrv = &taskSrvMock
	ctrl.CreateUserTask(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "CreateTask", 1)
}

func TestCreateUserTaskBindError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s/tasks", notUserID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPost,
		endPoint,
		strings.NewReader(`{"title":"Juanito's first task"}`),
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(notUserID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("CreateTask").Return(models.Task{}, nil)

	ctrl := StartUser()
	ctrl.taskSrv = &taskSrvMock
	ctrl.CreateUserTask(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "CreateTask", 0)
}

func TestCreateUserTaskServiceError(t *testing.T) {
	endPoint := fmt.Sprintf("/api/users/%s/tasks", userID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodPost,
		endPoint,
		strings.NewReader(`{"title":"Juanito's first task"}`),
	)

	c.SetParamNames(paramUserID)
	c.SetParamValues(userID)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("CreateTask").Return(models.Task{}, errUser)

	ctrl := StartUser()
	ctrl.taskSrv = &taskSrvMock
	ctrl.CreateUserTask(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	taskSrvMock.AssertNumberOfCalls(t, "CreateTask", 1)
}
