package task

import (
	"context"

	models "github.com/hjcalderon10/bunny-backend/model"
)

func New() service {
	return service{}
}

func (srv service) CreateTask(ctx context.Context, task models.Task) error {
	return nil
}
func (srv service) ReadTask(ctx context.Context, task *models.Task) error {
	return nil
}
func (srv service) UpdateTask(ctx context.Context, task *models.Task) error {
	return nil
}
func (srv service) DeleteTask(ctx context.Context, task models.Task) error {
	return nil
}
