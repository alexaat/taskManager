/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"taskManager/database"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",

	Run: func(cmd *cobra.Command, args []string) {

		var indexes []int
		for _, val := range args {
			index, err := strconv.Atoi(val)
			if err != nil {
				fmt.Printf("Cannot convert %s to int\n", string(val))
				os.Exit(1)
			} else {
				indexes = append(indexes, index)
			}
		}
		if len(indexes) > 0 {
			marksSuccessful, err := database.DoTasks(indexes)
			if err != nil {
				fmt.Println("Error while marking tasks as done", err)
				os.Exit(1)
			}
			fmt.Printf("Marked as comlete: %s\n", strings.Join(marksSuccessful, ","))
		} else {
			fmt.Println("No tasks to mark as complete")
		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
