package set

import (
	"fmt"
	"strconv"

	"github.com/ioanrobertrosu/todo/internal/task"
	"github.com/spf13/cobra"
)

var doneCommand = cobra.Command{
	Use: "done <id>",
	Short: "Set the selected task as [done]",
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

		task, err := taskList.Get(taskID)
		if err != nil {
			return err
		}

		task.Done()

		if err := taskList.Store(); err != nil {
			return err
		}

		fmt.Printf("task [%d] marketd as [done]", taskID)

		return nil
	},
}