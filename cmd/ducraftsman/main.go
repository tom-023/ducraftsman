package main

import (
	"github.com/tom-023/ducraftsman/pkg/cli"
	"os"
)

func main() {
	// ルートコマンドの作成と実行
	rootCmd := cli.NewRootCmd()

	// コマンドの実行
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
