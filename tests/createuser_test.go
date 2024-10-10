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

	// Create関数を呼び出し
	err := ducraftsman.Create(mockDBManager, "root", "password", "testdb", "localhost", "newuser", "newpassword", "ALL")

	// エラーが発生しないことを確認
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}

func TestCreateUser_FailOnConnect(t *testing.T) {
	// 接続失敗を模擬
	mockDBManager := &MockDBManager{
		ConnectFunc: func(rootUser, rootPassword, host, dbName string) (*sql.DB, error) {
			//return nil, errors.New("failed to connect")
			db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
			return db, nil
		},
		CreateUserFunc: func(db *sql.DB, username, password, privileges string) error {
			return errors.New("failed to create user")
		},
	}

	// Create関数を呼び出し
	err := ducraftsman.Create(mockDBManager, "root", "wrongpassword", "testdb", "localhost", "newuser", "newpassword", "ALL")

	// エラーが発生したことを確認
	if err == nil {
		t.Fatalf("expected error but got none")
	}

	// エラーメッセージの確認
	expectedError := "failed to create user: failed to create user"
	if err.Error() != expectedError {
		t.Errorf("expected %s but got %s", expectedError, err.Error())
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

	// Create関数を呼び出し
	err := ducraftsman.Create(mockDBManager, "root", "password", "testdb", "localhost", "newuser", "newpassword", "ALL")

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
