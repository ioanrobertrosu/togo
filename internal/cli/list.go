package cli

import (
	"fmt"

	"github.com/ioanrobertrosu/todo/internal/task"
	"github.com/spf13/cobra"
)

var listCommand = cobra.Command{
	Use: "list",
	Short: "Displays all the available tasks",

	RunE: func(cmd *cobra.Command, args []string) error {
		taskList := task.TaskList{}
		if err := taskList.Load(); err != nil {
			return err
		}

		fmt.Printf("%s", taskList.List())

		return nil
	},
}