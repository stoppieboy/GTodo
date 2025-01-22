package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Aliases: []string{"rm", "d"},
	Short: "Delete a task from the list",
	Long: "Delete a task from the list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Something went wrong: %v", err)
			return
		}
		err = Delete(id)
		if err != nil {
			fmt.Printf("Something went wrong: %v", err)
			return
		}
		fmt.Printf("Task Deleted: %v", args[0])
	},
}