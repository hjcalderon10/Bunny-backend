package user

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

func TestCreateUser(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestCreateUser"))
	clearDB()
	repo = New()
	defer clearDB()

	user := models.User{
		Name:     "userName",
		NickName: "Nickk",
		ImgURL:   "theImg.jpg",
	}

	err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)

	rows, err := repo.db.Raw(fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", user.Name))
	assert.Nil(t, err)
	defer rows.Close()

	userDB := models.User{}
	if rows.Next() {
		sqlstruct.Scan(&userDB, rows)
	}

	assert.Equal(t, user.Name, userDB.Name)
	assert.Equal(t, user.NickName, userDB.NickName)
	assert.Equal(t, user.ImgURL, userDB.ImgURL)
}

func TestCreateUserError(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestCreateUserError"))
	repo = New()
	bdMock := mocks.StorageRepoMock{}
	bdMock.On("Exec").Return(nil, errRepo)
	repo.db = &bdMock

	user := models.User{
		Name:     "userName",
		NickName: "Nickk",
		ImgURL:   "theImg.jpg",
	}

	err := repo.CreateUser(ctx, user)
	assert.Error(t, err)

	bdMock.AssertNumberOfCalls(t, "Exec", 1)
}

func TestGetAllUsers(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestGetAllUsers"))
	clearDB()
	repo = New()
	defer clearDB()

	users := []models.User{
		models.User{
			Name:     "userName",
			NickName: "Nickk",
			ImgURL:   "theImg.jpg",
		},
		models.User{
			Name:     "userName2",
			NickName: "NickName2",
			ImgURL:   "TheImg2.jpg",
		},
	}

	for _, user := range users {
		err := repo.CreateUser(ctx, user)
		assert.NoError(t, err)
	}

	usersDB, err := repo.GetAllUsers(ctx)
	assert.NoError(t, err)

	assert.Equal(t, len(users), len(usersDB))
}

func TestGetAllUsersError(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestGetAllUsersError"))
	repo = New()

	bdMock := mocks.StorageRepoMock{}
	bdMock.On("Raw").Return(nil, errRepo)
	repo.db = &bdMock

	users, err := repo.GetAllUsers(ctx)
	assert.Error(t, err)

	assert.Equal(t, 0, len(users))

	bdMock.AssertNumberOfCalls(t, "Raw", 1)
}

func TestUpdateUser(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestUpdateUser"))
	clearDB()
	repo = New()
	defer clearDB()

	user := models.User{
		Name: "My created user",
	}

	err := repo.CreateUser(ctx, user)

	fmt.Println(err)
	rows, err := repo.db.Raw(fmt.Sprintf("SELECT id FROM users WHERE name='%s'", user.Name))
	assert.NoError(t, err)
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&user.ID)
	}

	assert.NoError(t, rows.Err())
	assert.NotEqual(t, 0, user.ID)

	user.Name = "Updated name"
	err = repo.UpdateUser(ctx, user)
	assert.NoError(t, err)

	userDB, err := repo.ReadUser(ctx, user.ID)
	assert.NoError(t, err)

	assert.Equal(t, user.Name, userDB.Name)
	assert.Equal(t, user.ID, userDB.ID)
	assert.Equal(t, user.NickName, userDB.NickName)
	assert.Equal(t, user.ImgURL, userDB.ImgURL)
}

func TestUpdateUserError(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestUpdateUserError"))
	repo = New()
	bdMock := mocks.StorageRepoMock{}
	bdMock.On("Exec").Return(nil, errRepo)
	repo.db = &bdMock

	user := models.User{
		Name: "My created user",
	}

	err := repo.UpdateUser(ctx, user)
	assert.Error(t, err)
	bdMock.AssertNumberOfCalls(t, "Exec", 1)

}

func TestDeleteUser(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestDeleteUser"))
	clearDB()
	repo = New()
	defer clearDB()

	user := models.User{
		Name: "My created user",
	}

	err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)

	rows, err := repo.db.Raw(fmt.Sprintf("SELECT id FROM users WHERE name='%s'", user.Name))
	assert.NoError(t, err)
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&user.ID)
	}

	assert.NoError(t, rows.Err())
	assert.NotEqual(t, 0, user.ID)

	err = repo.DeleteUser(ctx, user.ID)
	assert.NoError(t, err)

	userDB, err := repo.ReadUser(ctx, user.ID)
	assert.NoError(t, err)

	assert.Equal(t, models.User{}, userDB)
}

func TestDeleteUserError(t *testing.T) {
	ctx := context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("TestDeleteUserError"))
	repo = New()
	bdMock := mocks.StorageRepoMock{}
	bdMock.On("Exec").Return(nil, errRepo)
	repo.db = &bdMock

	err := repo.DeleteUser(ctx, models.UserID(10))
	assert.Error(t, err)
	bdMock.AssertNumberOfCalls(t, "Exec", 1)

}
