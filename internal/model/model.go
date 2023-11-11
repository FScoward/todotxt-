package model

type Metadata map[string]string

type Task struct {
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
