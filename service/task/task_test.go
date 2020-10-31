package task

import (
	"context"
	"fmt"
	"testing"

	"github.com/hjcalderon10/bunny-backend/model"
	models "github.com/hjcalderon10/bunny-backend/model"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	mocks "github.com/hjcalderon10/bunny-backend/test/mocks/repository"
	"github.com/hjcalderon10/bunny-backend/util/logger"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTasksEmpty(t *testing.T) {
	tasks := []model.Task{}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("GetAllTasks").Return(tasks, nil)
	srv := New(&taskRepoMock)

	res, err := srv.GetAllTasks(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")))
	assert.NoError(t, err)
	assert.Equal(t, tasks, res)
	taskRepoMock.AssertNumberOfCalls(t, "GetAllTasks", 1)
}

func TestGetAllTasks(t *testing.T) {
	tasks := []model.Task{
		model.Task{
			ID:    uint16(10),
			Title: "Vero Snow",
		},
		model.Task{
			ID:    uint16(11),
			Title: "Ciri of Rivia",
		},
	}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("GetAllTasks").Return(tasks, nil)
	srv := New(&taskRepoMock)

	res, err := srv.GetAllTasks(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")))
	assert.NoError(t, err)
	assert.Equal(t, tasks, res)
	taskRepoMock.AssertNumberOfCalls(t, "GetAllTasks", 1)
}

func TestGetAllTasksRepoError(t *testing.T) {
	tasks := []model.Task{}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("GetAllTasks").Return(tasks, fmt.Errorf("the wild hunt"))
	srv := New(&taskRepoMock)

	res, err := srv.GetAllTasks(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")))
	assert.Error(t, err)
	assert.Equal(t, tasks, res)
	taskRepoMock.AssertNumberOfCalls(t, "GetAllTasks", 1)
}

func TestCreateTask(t *testing.T) {
	task := models.Task{
		ID: uint16(5),
	}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("CreateTask").Return(nil)
	srv := New(&taskRepoMock)

	err := srv.CreateTask(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), task)

	assert.NoError(t, err)
	taskRepoMock.AssertNumberOfCalls(t, "CreateTask", 1)

}

func TestCreateTaskRepoError(t *testing.T) {
	task := models.Task{
		ID: uint16(5),
	}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("CreateTask").Return(fmt.Errorf("theres no chance u can create that task"))
	srv := New(&taskRepoMock)

	err := srv.CreateTask(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), task)

	assert.Error(t, err)
	taskRepoMock.AssertNumberOfCalls(t, "CreateTask", 1)

}

func TestReadTask(t *testing.T) {
	task := models.Task{
		ID: uint16(20),
	}
	taskMock := models.Task{
		ID:    uint16(20),
		Title: "Ganon of the Night",
	}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("ReadTask").Return(taskMock, nil)
	srv := New(&taskRepoMock)

	err := srv.ReadTask(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), &task)
	assert.NoError(t, err)
	assert.Equal(t, taskMock, task)
	taskRepoMock.AssertNumberOfCalls(t, "ReadTask", 1)
}

func TestReadTaskRepoError(t *testing.T) {
	task := models.Task{
		ID: uint16(20),
	}
	taskMock := models.Task{
		ID:    uint16(20),
		Title: "Ganon of the Night",
	}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("ReadTask").Return(taskMock, fmt.Errorf("oh-oh, there's light in the dawn"))
	srv := New(&taskRepoMock)

	err := srv.ReadTask(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), &task)
	assert.Error(t, err)
	assert.NotEqual(t, taskMock, task)
	taskRepoMock.AssertNumberOfCalls(t, "ReadTask", 1)
}

func TestUpdateTask(t *testing.T) {
	task := models.Task{}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("UpdateTask").Return(nil)
	srv := New(&taskRepoMock)

	err := srv.UpdateTask(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), task)
	assert.NoError(t, err)
	taskRepoMock.AssertNumberOfCalls(t, "UpdateTask", 1)
}

func TestUpdateTaskRepoError(t *testing.T) {
	task := models.Task{}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("UpdateTask").Return(fmt.Errorf("err 404, task not found"))
	srv := New(&taskRepoMock)

	err := srv.UpdateTask(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), task)
	assert.Error(t, err)
	taskRepoMock.AssertNumberOfCalls(t, "UpdateTask", 1)
}

func TestDeleteTask(t *testing.T) {
	task := models.Task{
		ID: uint16(99),
	}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("DeleteTask").Return(nil)
	srv := New(&taskRepoMock)

	err := srv.DeleteTask(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), task)
	assert.NoError(t, err)
	taskRepoMock.AssertNumberOfCalls(t, "DeleteTask", 1)
}

func TestDeleteTaskRepoError(t *testing.T) {
	task := models.Task{
		ID: uint16(99),
	}
	taskRepoMock := mocks.TaskMock{}
	taskRepoMock.On("DeleteTask").Return(fmt.Errorf("once again, err 404, task not found"))
	srv := New(&taskRepoMock)

	err := srv.DeleteTask(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), task)
	assert.Error(t, err)
	taskRepoMock.AssertNumberOfCalls(t, "DeleteTask", 1)
}
