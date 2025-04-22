package cmd

import (
	"fmt"
	"os"

	"github.com/mihabgit/todo-cli/todo"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all TODO tasks",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := todo.Load(dataFile)
		if err != nil {
			fmt.Println("Error loading todos:", err)
			os.Exit(1)
		}
		for i, t := range todos {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("%d. [%s] %s\n", i+1, status, t.Text)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
