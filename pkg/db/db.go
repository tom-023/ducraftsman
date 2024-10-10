package db

import (
	"database/sql"
	"fmt"
)

type DBManager interface {
	Connect(rootUser, rootPassword, host, dbName string) (*sql.DB, error)
	CreateUser(db *sql.DB, username, password, privileges string) error
}

func NewDBManager(dbType string) (DBManager, error) {
	switch dbType {
	case "mysql":
		return &MySQLManager{}, nil
	case "postgresql":
		return &MySQLManager{}, nil
		//return &PostgreSQLManager{}, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s. Use 'mysql' or 'postgresql'", dbType)
	}
}
