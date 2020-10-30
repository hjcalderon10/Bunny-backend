package task

import (
	"context"

	models "github.com/hjcalderon10/bunny-backend/model"
)

type (
	IService interface {
		ReadTask(ctx context.Context, task *models.Task) error
		UpdateTask(ctx context.Context, task *models.Task) error
		DeleteTask(ctx context.Context, task models.Task) error
	}

	ITaskService interface {
		CreateTask(ctx context.Context, task models.Task) error
	}

	IRepo interface {
	}

	service struct{}
)
