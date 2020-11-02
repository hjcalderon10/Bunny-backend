package task

import (
	"context"
	"fmt"
	"strings"

	models "github.com/hjcalderon10/bunny-backend/model"
	"github.com/hjcalderon10/bunny-backend/repository/storage"
	"github.com/kisielk/sqlstruct"
)

var repo *taskRepository

type taskRepository struct {
	db storage.StorageRepository
}

func New() *taskRepository {
	if repo == nil {
		repo = &taskRepository{storage.NewPostgres()}
	}
	return repo
}

func (repo taskRepository) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	tasks := []models.Task{}
	rows, err := repo.db.Raw(get_all_tasks)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	var task models.Task
	for rows.Next() {
		task = models.Task{}
		sqlstruct.Scan(&task, rows)

		tasks = append(tasks, task)
	}

	err = rows.Err()

	return tasks, err
}

func (repo taskRepository) GetAllTaskStates(ctx context.Context) ([]models.TaskState, error) {
	taskStates := []models.TaskState{}
	rows, err := repo.db.Raw(get_all_task_states)
	if err != nil {
		return taskStates, err
	}
	defer rows.Close()

	var taskState models.TaskState
	for rows.Next() {
		taskState = models.TaskState{}
		sqlstruct.Scan(&taskState, rows)

		taskStates = append(taskStates, taskState)
	}

	err = rows.Err()

	return taskStates, err
}

func (repo taskRepository) CreateTask(ctx context.Context, task models.Task) (uint16, error) {
	var id uint16
	rows, err := repo.db.Raw(create_task, task.Title, task.Description, task.UserID)
	if err == nil {
		if rows.Next() {
			rows.Scan(&id)
		}
	}
	return id, err
}

func (repo taskRepository) ReadTask(ctx context.Context, taskID models.TaskID) (models.Task, error) {
	task := models.Task{}
	rows, err := repo.db.Raw(get_task, taskID)

	if err == nil {
		defer rows.Close()

		if rows.Next() {
			sqlstruct.Scan(&task, rows)
		}

		err = rows.Err()
	}

	return task, err
}

func (repo taskRepository) UpdateTask(ctx context.Context, task models.Task) error {
	str, err := repo.buildUpdateQuery(task)
	if err == nil {
		_, err = repo.db.Exec(fmt.Sprintf(update_task, str), task.ID)
	}
	return err
}

func (repo taskRepository) DeleteTask(ctx context.Context, taskID models.TaskID) error {
	_, err := repo.db.Exec(delete_task, taskID)
	return err
}

func (repo taskRepository) buildUpdateQuery(task models.Task) (string, error) {
	var strBuff strings.Builder

	if task.Title != "" {
		fmt.Fprintf(&strBuff, "title='%s', ", task.Title)
	}
	if task.Description != "" {
		fmt.Fprintf(&strBuff, "description='%s', ", task.Description)
	}
	if task.State != "" {
		fmt.Fprintf(&strBuff, "state='%s', ", task.State)
	}

	if strBuff.Len() == 0 {
		return "", fmt.Errorf("No properties to update")
	}
	return strBuff.String()[:strBuff.Len()-2], nil
}
