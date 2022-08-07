package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
	Long: `

	task is a CLI for managing your TODOs.
	
	Usage:
	task [command]

  	Available Commands:
		add         Add a new task to your TODO list
		do          Mark a task on your TODO list as complete
		list        List all of your incomplete tasks
  
  	Use "task [command] --help" for more information about a command.
  `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
