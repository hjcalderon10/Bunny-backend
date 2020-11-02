package task

import (
	"context"

	models "github.com/hjcalderon10/bunny-backend/model"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	errors "github.com/hjcalderon10/bunny-backend/util/error"
	"github.com/hjcalderon10/bunny-backend/util/logger"
)

func New(repo IRepo) service {
	return service{
		repo: repo,
	}
}

func (srv service) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	tasks, err := srv.repo.GetAllTasks(ctx)

	if err != nil {
		log.Errorf("[GetAllTasks:%s]", err)
		err = errors.InternalServerError
	}

	return tasks, err
}

func (srv service) GetAllTaskStates(ctx context.Context) ([]models.TaskState, error) {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	taskStates, err := srv.repo.GetAllTaskStates(ctx)

	if err != nil {
		log.Errorf("[GetAllTasks:%s]", err)
		err = errors.InternalServerError
	}

	return taskStates, err
}

func (srv service) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	id, err := srv.repo.CreateTask(ctx, task)

	if err != nil {
		log.Errorf("[CreateTask:%s]", err)
		err = errors.InternalServerError
	}

	return models.Task{ID: id}, err
}

func (srv service) ReadTask(ctx context.Context, task *models.Task) error {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	taskDB, err := srv.repo.ReadTask(ctx, task.ID)
	if err != nil {
		log.Errorf("[ReadTask:%s]", err)
		err = errors.InternalServerError

	} else {
		task.ID = taskDB.ID
		task.Title = taskDB.Title
		task.State = taskDB.State
		task.Description = taskDB.Description
		task.CreatedAt = taskDB.CreatedAt
		task.UserID = taskDB.UserID
	}

	return err
}

func (srv service) UpdateTask(ctx context.Context, task models.Task) error {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	err := srv.repo.UpdateTask(ctx, task)

	if err != nil {
		log.Errorf("[UpdateTask:%s]", err)
		err = errors.InternalServerError
	}

	return err
}
func (srv service) DeleteTask(ctx context.Context, task models.Task) error {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	err := srv.repo.DeleteTask(ctx, task.ID)

	if err != nil {
		log.Errorf("[DeleteTask:%s]", err)
		err = errors.InternalServerError
	}

	return err
}
