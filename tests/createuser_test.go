package tests

import (
	"bytes"
	"database/sql"
	"errors"
	"github.com/tom-023/ducraftsman/internal/cli"
	"testing"
)

func TestCreateUser_Success(t *testing.T) {
	mockDBManager := &MockDBManager{
		// モックされたConnectFuncを使用してエラーなしで接続成功を模擬
		ConnectFunc: func(rootUser, rootPassword, host, dbName string) (*sql.DB, error) {
			db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
			return db, nil
			//return &sql.DB{}, nil // 空のsql.DBを返す
		},
		// モックされたCreateUserFuncを使用してユーザ作成成功を模擬
		CreateUserFunc: func(db *sql.DB, username, password, privileges string) error {
			return nil
		},
	}

	cmd := cli.NewCreateUserCmd(mockDBManager)

	// フラグと引数を設定
	args := []string{
		"--rootuser", "root",
		"--rootpassword", "password",
		"--dbname", "testdb",
		"--dbhost", "localhost",
		"--username", "newuser",
		"--userpassword", "newpassword",
		"--privileges", "ALL",
	}
	cmd.SetArgs(args)

	// 出力をキャプチャ
	output := new(bytes.Buffer)
	cmd.SetOut(output)

	// コマンドの実行
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}

	// 出力の検証
	expectedOutput := "User newuser created successfully with password: newpassword\n"
	if output.String() != expectedOutput {
		t.Errorf("expected %s but got %s", expectedOutput, output.String())
	}
}

func TestCreateUser_FailOnConnect(t *testing.T) {
	mockDBManager := &MockDBManager{
		// モックされたConnectFuncを使用して接続失敗を模擬
		ConnectFunc: func(rootUser, rootPassword, host, dbName string) (*sql.DB, error) {
			db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
			return db, errors.New("failed to connect")
			//return nil, errors.New("failed to connect")
		},
		CreateUserFunc: func(db *sql.DB, username, password, privileges string) error {
			return nil
		},
	}

	cmd := cli.NewCreateUserCmd(mockDBManager)

	args := []string{
		"--rootuser", "root",
		"--rootpassword", "wrongpassword",
		"--dbname", "testdb",
		"--dbhost", "localhost",
		"--username", "newuser",
		"--userpassword", "newpassword",
		"--privileges", "ALL",
	}
	cmd.SetArgs(args)

	// 出力をキャプチャ
	output := new(bytes.Buffer)
	cmd.SetOut(output)

	// コマンドの実行
	err := cmd.Execute()
	if err == nil {
		t.Fatalf("expected an error but got none")
	}

	// エラーメッセージの検証
	expectedError := "failed to connect to database: failed to connect"
	if err.Error() != expectedError {
		t.Errorf("expected %s but got %s", expectedError, err.Error())
	}
}
