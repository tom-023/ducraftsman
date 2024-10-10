package cli

import (
	"github.com/spf13/cobra"
	"github.com/tom-023/ducraftsman"
	"github.com/tom-023/ducraftsman/pkg/db"
)

func NewRootCmd() *cobra.Command {
	var dbType, rootUser, rootPassword, dbName, dbHost, username, userPassword, privileges string

	// ルートコマンドの設定
	rootCmd := &cobra.Command{
		Use:   "ducraftsman",
		Short: "A tool for managing database users",
	}

	rootCmd.AddCommand(CreateUserCmd(&dbType, &rootUser, &rootPassword, &dbName, &dbHost, &username, &userPassword, &privileges))

	return rootCmd
}

// createコマンドにフラグを渡す形で実行する関数
func CreateUserCmd(dbType, rootUser, rootPassword, dbName, dbHost, username, userPassword, privileges *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new database user",
		RunE: func(cmd *cobra.Command, args []string) error {
			// db.NewDBManagerでDBManagerを選択
			dbManager, err := db.NewDBManager(*dbType)
			if err != nil {
				return err
			}
			// コマンド引数が正しく設定されているかを検証する処理などを行う

			return ducraftsman.Create(dbManager, *rootUser, *rootPassword, *dbName, *dbHost, *username, *userPassword, *privileges)
		},
	}
	// フラグの設定（CLIの引数を定義）
	cmd.Flags().StringVarP(dbType, "dbtype", "t", "mysql", "Database type (mysql or postgresql)")
	cmd.Flags().StringVarP(rootUser, "rootuser", "r", "", "Root user (required)")
	cmd.Flags().StringVarP(rootPassword, "rootpassword", "p", "", "Root password (required)")
	cmd.Flags().StringVarP(dbName, "dbname", "d", "", "Target database name (required)")
	cmd.Flags().StringVarP(dbHost, "dbhost", "H", "localhost", "Database host (default: localhost)")
	cmd.Flags().StringVarP(username, "username", "u", "", "Username to be created (required)")
	cmd.Flags().StringVarP(userPassword, "userpassword", "P", "", "Password for the new user (required)")
	cmd.Flags().StringVarP(privileges, "privileges", "g", "ALL", "Privileges to grant to the user (default: ALL)")

	// 必須フラグの設定（CLIの引数を定義）
	cmd.MarkFlagRequired("rootuser")
	cmd.MarkFlagRequired("rootpassword")
	cmd.MarkFlagRequired("dbname")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("userpassword")

	return cmd
}
