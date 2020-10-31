package service

import (
	"context"

	models "github.com/hjcalderon10/bunny-backend/model"
	"github.com/stretchr/testify/mock"
)

type TaskMock struct {
	mock.Mock
}

func (_m *TaskMock) CreateTask(ctx context.Context, task models.Task) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}

func (_m *TaskMock) ReadTask(ctx context.Context, task *models.Task) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}

func (_m *TaskMock) UpdateTask(ctx context.Context, task *models.Task) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}

func (_m *TaskMock) DeleteTask(ctx context.Context, task models.Task) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}
