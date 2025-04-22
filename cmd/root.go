package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var dataFile string

var rootCmd = &cobra.Command{
	Use: "todo-cli",
	Short: "TODO is a simple CLI TODO manager",
	Long: `A simple CLI application to manage your TODO list, build with Cobra in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'todo-cli help' to see available commands")
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&dataFile, "data", "d", "~/.todo.json", "data file to store todos")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
