package storage

import (
	"database/sql"
	"fmt"

	settings "github.com/hjcalderon10/bunny-backend/setting"
)

func NewPostgreStorage(settings settings.PostgresSettings) StorageRepository {
	conURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&connect_timeout=%d",
		settings.Username,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.Name,
		"disable",
		settings.DbPQTimeout)

	db, err := sql.Open("postgres", conURL)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(settings.MaxIdleConnections)
	db.SetMaxOpenConns(settings.MaxConnections)
	db.SetConnMaxLifetime(settings.ConnMaxLifeTime)

	return storageRepo{db}
}

func (s storageRepo) Raw(sql string, values ...interface{}) (*sql.Rows, error) {
	rows, err := s.db.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

type ExectStmt func(*sql.Tx) error

func (s storageRepo) DoTransaction(fnStmt ExectStmt) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	err = fnStmt(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s storageRepo) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, args...)
}
