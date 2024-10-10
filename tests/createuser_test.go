package tests

import (
	"database/sql"
	"errors"
	"github.com/tom-023/ducraftsman"
	"testing"
)

func TestCreateUser_Success(t *testing.T) {
	// モックのDBManagerを作成
	mockDBManager := &MockDBManager{
		ConnectFunc: func(rootUser, rootPassword, host, dbName string) (*sql.DB, error) {
			db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
			return db, nil
		},
		CreateUserFunc: func(db *sql.DB, username, password, privileges string) error {
			// 成功時はエラーなしでユーザ作成
			return nil
		},
	}

	rootUser := "root"
	rootPassword := "password"
	dbName := "testdb"
	dbHost := "localhost"
	username := "newuser"
	privileges := "ALL"

	// Create関数を呼び出し
	err := ducraftsman.Create(mockDBManager, rootUser, rootPassword, dbName, dbHost, username, privileges)

	// エラーが発生しないことを確認
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}

func TestCreateUser_FailOnCreateUser(t *testing.T) {
	// ユーザ作成失敗を模擬
	mockDBManager := &MockDBManager{
		ConnectFunc: func(rootUser, rootPassword, host, dbName string) (*sql.DB, error) {
			db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
			return db, nil
		},
		CreateUserFunc: func(db *sql.DB, username, password, privileges string) error {
			return errors.New("error test")
		},
	}

	rootUser := "root"
	rootPassword := "password"
	dbName := "testdb"
	dbHost := "localhost"
	username := "newuser"
	privileges := "ALL"

	// Create関数を呼び出し
	err := ducraftsman.Create(mockDBManager, rootUser, rootPassword, dbName, dbHost, username, privileges)

	// エラーが発生したことを確認
	if err == nil {
		t.Fatalf("expected error but got none")
	}

	// エラーメッセージの確認
	expectedError := "failed to create user: error test"
	if err.Error() != expectedError {
		t.Errorf("expected %s but got %s", expectedError, err.Error())
	}
}
