package task

import "fmt"

type TaskInput struct {
	Text string
}

func (ts TaskInput) ToTask(tl *TaskList) (*Task, error) {
	if ts.Text == "" {
		return nil, fmt.Errorf("task [Text] field cannot be empty")
	}

	task := Task{
		ID:        tl.GetNextTaskID(),
		Text:      ts.Text,
		Completed: false,
	}

	if task.ID <= 0 {
		return nil, fmt.Errorf("task [ID] field canno't be less or equal to 0")
	}

	return &task, nil
}