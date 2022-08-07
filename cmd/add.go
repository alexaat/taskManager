package cmd

import (
	"fmt"
	"os"
	"strings"
	"taskManager/database"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to our list",
	Run: func(cmd *cobra.Command, args []string) {

		task := strings.Join(args, " ")

		err := database.AddTask(task)

		if err != nil {
			fmt.Printf("Error while adding %v: %v\n", task, err)
			os.Exit(1)
		}

		fmt.Printf("%v was added to database\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
