package ducraftsman

import (
	"fmt"
	"github.com/tom-023/ducraftsman/pkg/db"
)

// Create関数は、新しいデータベースユーザを作成するビジネスロジックを担当
func Create(dbManager db.DBManager, rootUser, rootPassword, dbName, dbHost, username, userPassword, privileges string) error {
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

	// 成功メッセージを出力
	fmt.Printf("User %s created successfully with password: %s\n", username, userPassword)
	return nil
}
