package cmd

import (
	"fmt"
	"strings"

	"amru.in/cli/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Add called")
		task := strings.Join(args, " ")
		_, err := db.AddTask(task)

		if err != nil {
			fmt.Println("Unable to add task. Error occured: ", err)
		}
		fmt.Println("Task added to the list")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
