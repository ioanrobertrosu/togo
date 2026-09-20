package task

import (
	"testing"
)

func TestToTaskEmptyInput(t *testing.T) {
	tl := TaskList{Tasks: []Task{}}
	input := TaskInput{Text: ""}

	_, err := input.ToTask(&tl)
	if err == nil {
		t.Errorf("expected %q, got nil", "task [Text] field cannot be empty")
	}
}

func TestToTaskIdLessOrEqualToZero(t *testing.T) {
	errCount := 0
	tl := TaskList{
		Tasks: []Task{
			{
				ID: -2,
				Text: "task 1",
				Completed: true,
			},
		},
	}

	input := TaskInput{Text: "task 2"}

	for tl.Tasks[0].ID <= 2 {
		_, err := input.ToTask(&tl)

		if err != nil {
			errCount++
		}

		tl.Tasks[0].ID++
	}

	if errCount != 3 {
		t.Errorf("expected 3 errors, go %d", errCount)
	}
}