package repository

import (
	"database/sql"

	"github.com/hjcalderon10/bunny-backend/repository/storage"
	"github.com/stretchr/testify/mock"
)

type StorageRepoMock struct {
	mock.Mock
}

func (_m *StorageRepoMock) Raw(query string, values ...interface{}) (*sql.Rows, error) {
	args := _m.Called()

	if err := args.Get(1); err != nil {
		return nil, err.(error)
	}

	return args.Get(0).(*sql.Rows), nil
}

func (_m *StorageRepoMock) DoTransaction(fnStmt storage.ExectStmt) error {
	args := _m.Called()

	if err := args.Get(0); err != nil {
		return err.(error)
	}

	return nil
}

func (_m *StorageRepoMock) Exec(query string, argsParams ...interface{}) (sql.Result, error) {
	args := _m.Called()

	if err := args.Get(1); err != nil {
		return nil, err.(error)
	}

	return args.Get(0).(sql.Result), nil
}
