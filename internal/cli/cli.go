package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ducraftsman",
		Short: "ducraftsman is a tool to manage database users",
		Long:  `ducraftsman is a CLI tool that allows you to create and manage database users with ease.`,
	}

	return rootCmd
}
