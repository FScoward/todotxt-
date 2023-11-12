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
	// TODO 作成日時
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

func (task *Task) AddSubTask(targetTaskId ulid.ULID, description string) (Task, error) {

	newSubTask := Task{
		ID:          ulid.Make(),
		Status:      "TODO",
		Description: description,
	}

	task.SubTasks = append(task.SubTasks, newSubTask)
	return *task, nil
}
