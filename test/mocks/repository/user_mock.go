package repository

import (
	"context"

	models "github.com/hjcalderon10/bunny-backend/model"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (_m *UserMock) GetAllUsers(ctx context.Context) ([]models.User, error) {
	args := _m.Called()

	if err := args.Get(1); err != nil {
		return []models.User{}, err.(error)
	}

	return args.Get(0).([]models.User), nil
}

func (_m *UserMock) CreateUser(ctx context.Context, user models.User) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}

func (_m *UserMock) ReadUser(ctx context.Context, userID models.UserID) (models.User, error) {
	args := _m.Called()

	if err := args.Get(1); err != nil {
		return models.User{}, err.(error)
	}

	return args.Get(0).(models.User), nil
}

func (_m *UserMock) UpdateUser(ctx context.Context, user models.User) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}

func (_m *UserMock) DeleteUser(ctx context.Context, userID models.UserID) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}
