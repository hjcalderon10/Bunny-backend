package repository

import (
	"context"

	models "github.com/hjcalderon10/bunny-backend/model"
	"github.com/stretchr/testify/mock"
)

type TaskMock struct {
	mock.Mock
}

func (_m *TaskMock) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	args := _m.Called()

	if err := args.Get(1); err != nil {
		return []models.Task{}, err.(error)
	}

	return args.Get(0).([]models.Task), nil
}

func (_m *TaskMock) GetAllTaskStates(ctx context.Context) ([]models.TaskState, error) {
	args := _m.Called()

	if err := args.Get(1); err != nil {
		return []models.TaskState{}, err.(error)
	}

	return args.Get(0).([]models.TaskState), nil
}

func (_m *TaskMock) CreateTask(ctx context.Context, task models.Task) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}

func (_m *TaskMock) ReadTask(ctx context.Context, taskID models.TaskID) (models.Task, error) {
	args := _m.Called()

	if err := args.Get(1); err != nil {
		return models.Task{}, err.(error)
	}

	return args.Get(0).(models.Task), nil
}

func (_m *TaskMock) UpdateTask(ctx context.Context, task models.Task) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}

func (_m *TaskMock) DeleteTask(ctx context.Context, taskID models.TaskID) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}
