package cmd

import (
	"fmt"
	"os"

	"github.com/mihabgit/todo-cli/todo"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t := todo.Todo{Text: args[0]}
		todos, err := todo.Load(dataFile)
		if err != nil {
			fmt.Println("Error loading todos:", err)
			os.Exit(1)
		}
		todos = append(todos, t)
		if err := todo.Save(dataFile, todos); err != nil {
			fmt.Println("Error saving todos:", err)
			os.Exit(1)
		}
		fmt.Println("Added TODO:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
