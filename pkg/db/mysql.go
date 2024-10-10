package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLManager struct{}

func (m *MySQLManager) Connect(rootUser, rootPassword, host, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", rootUser, rootPassword, host, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (m *MySQLManager) CreateUser(db *sql.DB, username, password, privileges string) error {
	query := fmt.Sprintf("CREATE USER '%s'@'%%' IDENTIFIED BY '%s';", username, password)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	grantQuery := fmt.Sprintf("GRANT %s ON *.* TO '%s'@'%%';", privileges, username)
	_, err = db.Exec(grantQuery)
	return err
}
