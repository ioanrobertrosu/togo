package task

import (
	"testing"
)

func TestList(t *testing.T) {
	tests := []struct{
		name		string
		list		TaskList
		expected	string
	} {
		{
			name: "empty list",
			list: TaskList{
				Tasks: []Task{},
			},
			expected: "No tasks available",
		},
		{
            name: "undone task",
            list: TaskList{
                Tasks: []Task{
                    {
                        ID:        1,
                        Text:      "Learn Go testing",
                        Completed: false,
                    },
                },
            },
            expected: "[1] Learn Go testing -> undone\n",
        },
        {
            name: "done task",
            list: TaskList{
                Tasks: []Task{
                    {
                        ID:        1,
                        Text:      "Learn Go testing",
                        Completed: true,
                    },
                },
            },
			expected: "[1] Learn Go testing -> done\n",
		},
		{
			name: "multiple tasks",
			list: TaskList{
				Tasks: []Task{
					{
						ID: 1,
						Text: "task 1",
						Completed: false,
					},
					{
						ID: 2,
						Text: "task 2",
						Completed: true,
					},
					{
						ID: 3,
						Text: "task 3",
						Completed: false,
					},
				},
			},
			expected: "[1] task 1 -> undone\n" +
			"[2] task 2 -> done\n" +
			"[3] task 3 -> undone\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			result := tt.list.List()

			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct{
		name		string
		list		TaskList
		id			int
		expected	string 
	}{
		{
			name: "task id is less or equal to zero",
			list: TaskList{
				Tasks: []Task{
					{
						ID: 1,
						Text: "task 0",
						Completed: false,
					},
				},
			},
			id: 0,
			expected: "task [ID] field can't be less or equal to 0",
		},
		{
			name: "empty task list",
			list: TaskList{
				Tasks: []Task{},
			},
			id: 1,
			expected: "task list is empty",
		},
		{
			name: "task does not exist",
			list: TaskList{
				Tasks: []Task{
					{
						ID: 1,
						Text: "task 1",
						Completed: true,
					},
				},
			},
			id: 2,
			expected: "task [2] doesn't exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			_, err := tt.list.Get(tt.id)

			if err == nil {
				t.Fatalf("expected error %q, got nil", tt.expected)
			}

			if err.Error() != tt.expected {
				t.Errorf("expected error %q, got %q", tt.expected, err.Error())
			}
		})
	}
}