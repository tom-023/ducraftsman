package db

import (
	"database/sql"
)

// DBManagerインターフェース
type DBManager interface {
	Connect(rootUser, rootPassword, host, dbName string) (*sql.DB, error)
	CreateUser(db *sql.DB, username, password, privileges string) error
}
