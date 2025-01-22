package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"a"},
	Short: "Add a task to the list",
	Long: "Add a task to the list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := Add(args[0])
		if err != nil {
			fmt.Printf("Something went wrong: %v", err)
			return
		}
		fmt.Printf("Task Added: %v", args[0])
	},
}