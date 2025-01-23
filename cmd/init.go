package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

 var initCmd = &cobra.Command{
	Use: "init",
	Aliases: []string{"i"},
	Short: "Initializes the database",
	Long: "Initializes the database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing the database")
		err := Init()
		if err != nil {
			fmt.Println("Error initializing the database: ", err)
			return
		}
		fmt.Println("Database initialized successfully")
	},
}
