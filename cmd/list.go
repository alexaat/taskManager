package cmd

import (
	"fmt"
	"os"
	"taskManager/database"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of our incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := database.GetTasks(database.Uncompleted)
		if err != nil {
			fmt.Println("Error while getting tasks", err)
			os.Exit(1)
		}

		index := 1

		if len(tasks) <= 0 {
			fmt.Println("List is empty")
		} else {
			for _, v := range tasks {
				fmt.Printf("#%d. %s\n", index, v)
				index++
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
