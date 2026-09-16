package cli

import (
	"fmt"

	"github.com/ioanrobertrosu/todo/internal/task"
	"github.com/spf13/cobra"
)

var createCommand = cobra.Command{
	Use:   "create <text>",
	Short: "Create a new task",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		taskText := args[0]

		taskList := task.TaskList{}
		if err := taskList.Load(); err != nil {
			return err
		}

		TaskInput := task.TaskInput{Text: taskText}

		task, err := TaskInput.ToTask(&taskList)
		if err != nil {
			return err
		}

		taskList.Add(*task)
		if err := taskList.Store(); err != nil {
			return err
		}

		fmt.Printf("task [%d] created succesfully\n", task.ID)

		return nil
	},
}
