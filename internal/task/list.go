package task

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ioanrobertrosu/todo/internal/utils"
)

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func (tl TaskList) List() string {
	if len(tl.Tasks) == 0 {
		return "No tasks available"
	}

	var list strings.Builder

	for _, task := range tl.Tasks {
		state := "done"

		if !task.Completed {
			state = "undone"
		}
		list.WriteString("[")
		list.WriteString(strconv.Itoa(task.ID))
		list.WriteString("] ")
		list.WriteString(task.Text)
		list.WriteString(" -> ")
		list.WriteString(state)
		list.WriteString("\n")
	}

	return list.String()
}

func (tl TaskList) Get(id int) (*Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("task [ID] field can't be less or equal to 0")
	}

	if len(tl.Tasks) == 0 {
		return nil, fmt.Errorf("task list is empty")
	}

	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			return &tl.Tasks[i], nil
		}
	}

	return nil, fmt.Errorf("task [%d] doesn't exist", id)
}

func (tl *TaskList) Remove(id int) error {
	if id <= 0 {
		return fmt.Errorf("task id can't be less or equal to 0")
	}

	if len(tl.Tasks) == 0 {
		return fmt.Errorf("task list is empty")
	}

	idxToErase := -1

	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			idxToErase = i
			break
		}
	}

	if idxToErase == -1 {
		return fmt.Errorf("task [%d] not found", id)
	}

	tl.Tasks = append(tl.Tasks[:idxToErase], tl.Tasks[idxToErase+1:]...)

	return nil
}

func (tl *TaskList) Edit(id int, input TaskInput) error {
	taskToEdit, err := tl.Get(id)
	if err != nil {
		return err
	}

	if input.Text == "" {
		return fmt.Errorf("task [Text] field cannont be an empty string")
	}

	taskToEdit.Text = input.Text

	return nil
}

func (tl *TaskList) Add(t Task) {
	if tl.Tasks == nil {
		tl.Tasks = []Task{}
	}

	tl.Tasks = append(tl.Tasks, t)
}

func (tl TaskList) Store() error {
	file, err := utils.EnsureFileExists("tasks.json", true)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "     ")
	if err := encoder.Encode(tl); err != nil {
		return fmt.Errorf("couldn't encode JSON: %w", err)
	}

	return nil
}

func (tl *TaskList) Load() error {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			tl.Tasks = []Task{}
			return nil
		}

		return fmt.Errorf("couldn't open file: %w", err)
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return fmt.Errorf("couldn't check file status: %w", err)
	}

	if stat.Size() == 0 {
		tl.Tasks = []Task{}
		return nil
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(tl); err != nil {
		return fmt.Errorf("couldn't decode the information: %w", err)
	}

	return nil
}

func (tl TaskList) GetNextTaskID() int {
	if len(tl.Tasks) <= 0 {
		return 1
	}

	maxTaskID := tl.Tasks[0].ID

	for _, task := range tl.Tasks {
		if task.ID > maxTaskID {
			maxTaskID = task.ID
		}
	}

	return maxTaskID + 1
}
