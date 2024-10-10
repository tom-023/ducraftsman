package ducraftsman

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/tom-023/ducraftsman/pkg/db"
)

const PasswordLength = 12

func Create(dbManager db.DBManager, rootUser, rootPassword, dbName, dbHost, username, privileges string) error {
	// パスワードを内部で生成
	userPassword, err := generateRandomPassword(PasswordLength)
	if err != nil {
		return fmt.Errorf("failed to generate password: %v", err)
	}
	// DB接続
	dbConn, err := dbManager.Connect(rootUser, rootPassword, dbHost, dbName)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if dbConn != nil {
		defer dbConn.Close()
	}

	// ユーザ作成
	err = dbManager.CreateUser(dbConn, username, userPassword, privileges)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	fmt.Printf("User %s created successfully with password: %s\n", username, userPassword)
	return nil
}

func generateRandomPassword(n int) (string, error) {
	// nバイトのランダムなバイト列を生成
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate password: %v", err)
	}
	// Base64エンコードしてパスワードとして使用
	return base64.URLEncoding.EncodeToString(bytes)[:n], nil
}
