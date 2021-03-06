package cmd

import (
	"fmt"

	"amru.in/cli/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the to-do tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("list called")
		taskList, err := db.ListTaskItems()
		if err != nil {
			fmt.Println("Some error occured", err)
			return
		}

		if len(taskList) == 0 {
			fmt.Println("No tasks exist")
			return
		}
		fmt.Println("All tasks: ")

		for i, task := range taskList {
			fmt.Println((i + 1), task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
