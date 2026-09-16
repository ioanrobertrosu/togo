package task

type Task struct {
	ID   		int		`json:"id"`
	Text 		string	`json:"text"`
	Completed 	bool	`json:"completed"`
}

func (t *Task) Done() {
	t.Completed = true
}

func (t *Task) Undone() {
	t.Completed = false
}