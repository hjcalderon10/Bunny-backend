package task

import (
	"context"

	models "github.com/hjcalderon10/bunny-backend/model"
)

type (
	IService interface {
		GetAllTasks(ctx context.Context) ([]models.Task, error)
		ReadTask(ctx context.Context, task *models.Task) error
		UpdateTask(ctx context.Context, task models.Task) error
		DeleteTask(ctx context.Context, task models.Task) error
	}

	ITaskService interface {
		CreateTask(ctx context.Context, task models.Task) error
	}

	IRepo interface {
		GetAllTasks(ctx context.Context) ([]models.Task, error)
		CreateTask(ctx context.Context, task models.Task) error
		ReadTask(ctx context.Context, taskID models.TaskID) (models.Task, error)
		UpdateTask(ctx context.Context, task models.Task) error
		DeleteTask(ctx context.Context, taskID models.TaskID) error
	}

	service struct {
		repo IRepo
	}
)
