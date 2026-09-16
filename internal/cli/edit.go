package cli

import (
	"fmt"
	"strconv"

	"github.com/ioanrobertrosu/todo/internal/task"
	"github.com/spf13/cobra"
)

var editCommand = cobra.Command{
	Use: "edit <id>",
	Short: "Edit the selected task",
	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		input := task.TaskInput{Text: args[1]}

		taskList := task.TaskList{}
		if err := taskList.Load(); err != nil {
			return err
		}

		if err := taskList.Edit(taskID, input); err != nil {
			return err
		}

		if err := taskList.Store(); err != nil {
			return err
		}

		fmt.Printf("task [%d] edited sucessfuly\n", taskID)

		return nil
	},
}