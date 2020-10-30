package storage

import (
	"database/sql"
	"time"
)

type StorageRepository interface {
	Raw(sql string, values ...interface{}) (*sql.Rows, error)
	DoTransaction(fnStmt ExectStmt) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type storageRepo struct {
	db *sql.DB
}

type DocDbSetting struct {
	ConnString      string
	DbName          string
	Timeout         time.Duration
	MaxConnIdleTime int
	MaxConnections  int
}

type RedisSettings struct {
	Host        string
	Port        string
	DialTimeout time.Duration
	Timeout     time.Duration
	PoolSize    int
}
