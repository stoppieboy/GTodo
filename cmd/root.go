package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "GTodo",
		Short: "A CLI TODO Application",
		Long: "Track your tasks with local task persistence",
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(listCmd)
}

func Execute() {
	rootCmd.Execute()
}