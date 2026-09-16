package cli

import (
	"fmt"
	"strconv"

	"github.com/ioanrobertrosu/todo/internal/task"
	"github.com/spf13/cobra"
)

var deleteCommand = cobra.Command{
	Use: "delete <id>",
	Short: "Delete a task",
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		taskList := task.TaskList{}
		if err := taskList.Load(); err != nil {
			return err
		}

		if err := taskList.Remove(taskID); err != nil {
			return err
		}

		if err := taskList.Store(); err != nil {
			return err
		}

		fmt.Printf("task [%d] deleted successfuly\n", taskID)

		return nil
	},
}