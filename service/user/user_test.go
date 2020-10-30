package user

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

func TestGetAllUsersEmpty(t *testing.T) {
	users := []model.User{}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("GetAllUsers").Return(users, nil)
	srv := New(&userRepoMock)

	res, err := srv.GetAllUsers(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")))
	assert.NoError(t, err)
	assert.Equal(t, users, res)
	userRepoMock.AssertNumberOfCalls(t, "GetAllUsers", 1)
}

func TestGetAllUsers(t *testing.T) {
	users := []model.User{
		model.User{
			ID:   uint16(10),
			Name: "Vero Snow",
		},
		model.User{
			ID:   uint16(11),
			Name: "Ciri of Rivia",
		},
	}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("GetAllUsers").Return(users, nil)
	srv := New(&userRepoMock)

	res, err := srv.GetAllUsers(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")))
	assert.NoError(t, err)
	assert.Equal(t, users, res)
	userRepoMock.AssertNumberOfCalls(t, "GetAllUsers", 1)
}

func TestGetAllUsersRepoError(t *testing.T) {
	users := []model.User{}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("GetAllUsers").Return(users, fmt.Errorf("the wild hunt"))
	srv := New(&userRepoMock)

	res, err := srv.GetAllUsers(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")))
	assert.Error(t, err)
	assert.Equal(t, users, res)
	userRepoMock.AssertNumberOfCalls(t, "GetAllUsers", 1)
}

func TestCreateUser(t *testing.T) {
	user := models.User{
		ID: uint16(5),
	}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("CreateUser").Return(nil)
	srv := New(&userRepoMock)

	err := srv.CreateUser(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), user)

	assert.NoError(t, err)
	userRepoMock.AssertNumberOfCalls(t, "CreateUser", 1)

}

func TestCreateUserRepoError(t *testing.T) {
	user := models.User{
		ID: uint16(5),
	}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("CreateUser").Return(fmt.Errorf("theres no chance u can create that user"))
	srv := New(&userRepoMock)

	err := srv.CreateUser(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), user)

	assert.Error(t, err)
	userRepoMock.AssertNumberOfCalls(t, "CreateUser", 1)

}

func TestReadUser(t *testing.T) {
	user := models.User{
		ID: uint16(20),
	}
	userMock := models.User{
		ID:       uint16(20),
		Name:     "Ganon of the Night",
		NickName: "Pretty bad boy",
	}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("ReadUser").Return(userMock, nil)
	srv := New(&userRepoMock)

	err := srv.ReadUser(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), &user)
	assert.NoError(t, err)
	assert.Equal(t, userMock, user)
	userRepoMock.AssertNumberOfCalls(t, "ReadUser", 1)
}

func TestReadUserRepoError(t *testing.T) {
	user := models.User{
		ID: uint16(20),
	}
	userMock := models.User{
		ID:       uint16(20),
		Name:     "Ganon of the Night",
		NickName: "Pretty bad boy",
	}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("ReadUser").Return(userMock, fmt.Errorf("oh-oh, there's light in the dawn"))
	srv := New(&userRepoMock)

	err := srv.ReadUser(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), &user)
	assert.Error(t, err)
	assert.NotEqual(t, userMock, user)
	userRepoMock.AssertNumberOfCalls(t, "ReadUser", 1)
}

func TestUpdateUser(t *testing.T) {
	user := models.User{}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("UpdateUser").Return(nil)
	srv := New(&userRepoMock)

	err := srv.UpdateUser(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), user)
	assert.NoError(t, err)
	userRepoMock.AssertNumberOfCalls(t, "UpdateUser", 1)
}

func TestUpdateUserRepoError(t *testing.T) {
	user := models.User{}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("UpdateUser").Return(fmt.Errorf("err 404, user not found"))
	srv := New(&userRepoMock)

	err := srv.UpdateUser(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), user)
	assert.Error(t, err)
	userRepoMock.AssertNumberOfCalls(t, "UpdateUser", 1)
}

func TestDeleteUser(t *testing.T) {
	user := models.User{
		ID: uint16(99),
	}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("DeleteUser").Return(nil)
	srv := New(&userRepoMock)

	err := srv.DeleteUser(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), user)
	assert.NoError(t, err)
	userRepoMock.AssertNumberOfCalls(t, "DeleteUser", 1)
}

func TestDeleteUserRepoError(t *testing.T) {
	user := models.User{
		ID: uint16(99),
	}
	userRepoMock := mocks.UserMock{}
	userRepoMock.On("DeleteUser").Return(fmt.Errorf("once again, err 404, user not found"))
	srv := New(&userRepoMock)

	err := srv.DeleteUser(context.WithValue(context.Background(), settings.Commons.LogKey, logger.New("-")), user)
	assert.Error(t, err)
	userRepoMock.AssertNumberOfCalls(t, "DeleteUser", 1)
}
