package cmd

import (
    "os"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "GTodo",
		Short: "A CLI TODO Application",
		Long: "Track your tasks with local task persistence",
	}
	filePath = os.Getenv("USERPROFILE")+"\\GTodo\\tasks.csv"
	tempFilePath = os.Getenv("USERPROFILE")+"\\GTodo\\tasks.temp.csv"
)

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(listCmd)
}

func Execute() {
	rootCmd.Execute()
}
