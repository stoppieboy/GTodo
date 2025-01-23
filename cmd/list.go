package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"ls", "l"},
	Short: "Lists all the tasks",
	Long: "Lists all the tasks in the list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := List()
		if err != nil {
			fmt.Printf("Something went wrong: %v", err)
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.TabIndent)
		// fmt.Println("All Tasks ----------------")
		fmt.Fprintf(w, "No.\tTask\tStatus\tCreated At\n")
		fmt.Fprintf(w, "---\t----\t------\t----------\n")

		for i := range tasks {
			fmt.Fprintf(w, "%s", tasks[i])
		}
		w.Flush()
	},
}