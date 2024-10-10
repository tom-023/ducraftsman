package main

import (
	"github.com/tom-023/ducraftsman/internal/cli"
	"os"
)

func main() {
	rootCmd := cli.NewRootCmd()
	rootCmd.AddCommand(cli.NewCreateUserCmd()) // ユーザ作成コマンドを追加

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
