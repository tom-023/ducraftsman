package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tom-023/ducraftsman/pkg/db"
)

// NewCreateUserCmd ユーザ作成コマンド
func NewCreateUserCmd(dbManager db.DBManager) *cobra.Command {
	var rootUser, rootPassword, dbName, dbHost, username, userPassword, privileges string

	cmd := &cobra.Command{
		Use:   "createuser",
		Short: "Create a new database user",
		RunE: func(cmd *cobra.Command, args []string) error {
			// ここでDB接続とユーザ作成ロジックを呼び出す
			dbConn, err := dbManager.Connect(rootUser, rootPassword, dbHost, dbName)
			if err != nil {
				return fmt.Errorf("failed to connect to database: %v", err)
			}
			// dbConnがnilでない場合のみCloseを実行
			if dbConn != nil {
				defer dbConn.Close()
			}

			// ユーザ作成
			err = dbManager.CreateUser(dbConn, username, userPassword, privileges)
			if err != nil {
				return fmt.Errorf("failed to create user: %v", err)
			}

			// 標準出力に結果を出力
			fmt.Fprintf(cmd.OutOrStdout(), "User %s created successfully with password: %s\n", username, userPassword)
			return nil
		},
	}

	cmd.Flags().StringVarP(&rootUser, "rootuser", "r", "", "Root user (required)")
	cmd.Flags().StringVarP(&rootPassword, "rootpassword", "p", "", "Root password (required)")
	cmd.Flags().StringVarP(&dbName, "dbname", "d", "", "Target database name (required)")
	cmd.Flags().StringVarP(&dbHost, "dbhost", "H", "localhost", "Database host (default: localhost)")
	cmd.Flags().StringVarP(&username, "username", "u", "", "Username to be created (required)")
	cmd.Flags().StringVarP(&userPassword, "userpassword", "P", "", "Password for the new user (required)")
	cmd.Flags().StringVarP(&privileges, "privileges", "g", "ALL", "Privileges to grant to the user (default: ALL)")

	cmd.MarkFlagRequired("rootuser")
	cmd.MarkFlagRequired("rootpassword")
	cmd.MarkFlagRequired("dbname")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("userpassword")

	return cmd
}
