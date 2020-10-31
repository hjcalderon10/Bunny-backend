package task

import (
	"context"
	"fmt"
	"testing"

	models "github.com/hjcalderon10/bunny-backend/model"
	"github.com/hjcalderon10/bunny-backend/repository/storage"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	mocks "github.com/hjcalderon10/bunny-backend/test/mocks/repository"
	"github.com/hjcalderon10/bunny-backend/util/logger"
	"github.com/kisielk/sqlstruct"
	"github.com/stretchr/testify/assert"
)

var errRepo = fmt.Errorf("Something went wrong")

func clearDB() {
	repo = New()
	repo.db = storage.NewPostgres()
	_, err := repo.db.Raw("TRUNCATE TABLE users CASCADE;")
	if err != nil {
		panic(err)
	}
}

func TestCreateTask(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestCreateTask"))
	clearDB()
	repo = New()

	userID := models.UserID(1)

	_, err := repo.db.Raw(fmt.Sprintf("INSERT INTO users (id, name) VALUES (%d, 'MyFavUser');", userID))
	assert.NoError(t, err)

	task := models.Task{
		Title:  "taskTitle",
		UserID: userID,
	}

	err = repo.CreateTask(ctx, task)
	assert.NoError(t, err)

	rows, err := repo.db.Raw(fmt.Sprintf("SELECT * FROM tasks WHERE title = '%s'", task.Title))
	assert.Nil(t, err)
	defer rows.Close()

	taskDB := models.Task{}
	if rows.Next() {
		sqlstruct.Scan(&taskDB, rows)
	}

	assert.Equal(t, task.Title, taskDB.Title)
}

func TestCreateTaskError(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestCreateTaskError"))
	repo = New()
	bdMock := mocks.StorageRepoMock{}
	bdMock.On("Exec").Return(nil, errRepo)
	repo.db = &bdMock

	task := models.Task{
		Title: "taskTitle",
	}

	err := repo.CreateTask(ctx, task)
	assert.Error(t, err)

	bdMock.AssertNumberOfCalls(t, "Exec", 1)
}

func TestGetAllTasks(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestGetAllTasks"))
	clearDB()
	repo = New()

	userID := models.UserID(1)

	_, err := repo.db.Raw(fmt.Sprintf("INSERT INTO users (id, name) VALUES (%d, 'MyFavUser');", userID))
	assert.NoError(t, err)

	tasks := []models.Task{
		models.Task{
			UserID: userID,
			Title:  "taskTitle",
		},
		models.Task{
			UserID: userID,
			Title:  "taskTitle2",
		},
	}

	for _, task := range tasks {
		err := repo.CreateTask(ctx, task)
		assert.NoError(t, err)
	}

	tasksDB, err := repo.GetAllTasks(ctx)
	assert.NoError(t, err)

	assert.Equal(t, len(tasks), len(tasksDB))
}

func TestGetAllTasksError(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestGetAllTasksError"))
	repo = New()

	bdMock := mocks.StorageRepoMock{}
	bdMock.On("Raw").Return(nil, errRepo)
	repo.db = &bdMock

	tasks, err := repo.GetAllTasks(ctx)
	assert.Error(t, err)

	assert.Equal(t, 0, len(tasks))

	bdMock.AssertNumberOfCalls(t, "Raw", 1)
}

func TestUpdateTask(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestUpdateTask"))
	clearDB()
	repo = New()

	userID := models.UserID(1)

	_, err := repo.db.Raw(fmt.Sprintf("INSERT INTO users (id, name) VALUES (%d, 'MyFavUser');", userID))
	assert.NoError(t, err)

	task := models.Task{
		Title:  "My created task",
		UserID: userID,
	}

	err = repo.CreateTask(ctx, task)

	fmt.Println(err)
	rows, err := repo.db.Raw(fmt.Sprintf("SELECT id FROM tasks WHERE title='%s'", task.Title))
	assert.NoError(t, err)
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&task.ID)
	}

	assert.NoError(t, rows.Err())
	assert.NotEqual(t, 0, task.ID)

	task.Title = "Updated title"
	err = repo.UpdateTask(ctx, task)
	assert.NoError(t, err)

	taskDB, err := repo.ReadTask(ctx, task.ID)
	assert.NoError(t, err)

	assert.Equal(t, task.Title, taskDB.Title)
	assert.Equal(t, task.ID, taskDB.ID)
}

func TestUpdateTaskError(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestUpdateTaskError"))
	repo = New()
	bdMock := mocks.StorageRepoMock{}
	bdMock.On("Exec").Return(nil, errRepo)
	repo.db = &bdMock

	task := models.Task{
		Title: "My created task",
	}

	err := repo.UpdateTask(ctx, task)
	assert.Error(t, err)
	bdMock.AssertNumberOfCalls(t, "Exec", 1)

}

func TestDeleteTask(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestDeleteTask"))
	clearDB()
	repo = New()

	userID := models.UserID(1)

	_, err := repo.db.Raw(fmt.Sprintf("INSERT INTO users (id, name) VALUES (%d, 'MyFavUser');", userID))
	assert.NoError(t, err)

	task := models.Task{
		Title:  "My created task",
		UserID: userID,
	}

	err = repo.CreateTask(ctx, task)
	assert.NoError(t, err)

	rows, err := repo.db.Raw(fmt.Sprintf("SELECT id FROM tasks WHERE title='%s'", task.Title))
	assert.NoError(t, err)
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&task.ID)
	}

	assert.NoError(t, rows.Err())
	assert.NotEqual(t, 0, task.ID)

	err = repo.DeleteTask(ctx, task.ID)
	assert.NoError(t, err)

	taskDB, err := repo.ReadTask(ctx, task.ID)
	assert.NoError(t, err)

	assert.Equal(t, models.Task{}, taskDB)
}

func TestDeleteTaskError(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestDeleteTaskError"))
	repo = New()
	bdMock := mocks.StorageRepoMock{}
	bdMock.On("Exec").Return(nil, errRepo)
	repo.db = &bdMock

	err := repo.DeleteTask(ctx, models.TaskID(10))
	assert.Error(t, err)
	bdMock.AssertNumberOfCalls(t, "Exec", 1)

}
