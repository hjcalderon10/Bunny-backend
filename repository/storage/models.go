package storage

import (
	"database/sql"
)

type StorageRepository interface {
	Raw(sql string, values ...interface{}) (*sql.Rows, error)
	DoTransaction(fnStmt ExectStmt) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type storageRepo struct {
	db *sql.DB
}
