package model

import (
	"github.com/oklog/ulid/v2"
)

type Metadata map[string]string

type Task struct {
	ID          ulid.ULID
	Description string
	Status      string
	SubTasks    []Task
	Details     map[string]string
}

type Document struct {
	Tasks    []Task
	Meetings []string
	Logs     []string
	// その他のフィールド
}

type TaskList []Task

func (tasklist *TaskList) AddTask(ulid ulid.ULID, description string) (TaskList, error) {
	*tasklist = append(*tasklist, Task{
		ID:          ulid,
		Status:      "TODO",
		Description: description,
	})
	return *tasklist, nil
}
