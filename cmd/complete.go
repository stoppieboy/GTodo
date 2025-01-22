package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Aliases: []string{"c"},
	Short: "Mark a task as completed",
	Long: "Mark a task as completed",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Something went wrong: %v", err)
			return
		}
		err = Complete(id)
		if err != nil {
			fmt.Printf("Something went wrong: %v", err)
			return
		}
		fmt.Printf("Task Completed: %v", args[0])
	},
}