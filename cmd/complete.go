package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mihabgit/todo-cli/todo"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task number]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idx, err := strconv.Atoi(args[0])
		if err != nil || idx < 1 {
			fmt.Println("Please provide a valid task number")
			os.Exit(1)
		}
		todos, err := todo.Load(dataFile)
		if err != nil {
			fmt.Println("Error loading todos:", err)
			os.Exit(1)
		}
		if idx > len(todos) {
			fmt.Println("Task number out of range")
			os.Exit(1)
		}
		todos[idx-1].Done = true
		if err := todo.Save(dataFile, todos); err != nil {
			fmt.Println("Error saving todos:", err)
			os.Exit(1)
		}
		fmt.Printf("Marked TODO #%d as completed\n", idx)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
