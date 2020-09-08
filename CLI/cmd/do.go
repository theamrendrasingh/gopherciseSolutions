package cmd

import (
	"fmt"
	"strconv"

	"amru.in/cli/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks the to-do tasks as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var idx []int

		for _, val := range args {
			id, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(val, " is not a valid task id")
			} else {
				idx = append(idx, id)
			}
		}

		tasks, err := db.ListTaskItems()

		if err != nil {
			fmt.Println("Some error occured")
			return
		}

		for _, val := range idx {
			db.DeleteItem(tasks[val-1].Id)
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
