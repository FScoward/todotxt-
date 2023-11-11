package parser

import (
	"reflect"
	"strings"
	"testing"
	"todotxtplus/internal/model"

	"github.com/google/go-cmp/cmp"
)

func TestParseMetadata(t *testing.T) {
	lines := []string{
		"---",
		"projectName: ProjectX",
		"id: xxxxxx",
		"---",
	}
	expected := model.Metadata{
		"projectName": "ProjectX",
		"id":          "xxxxxx",
	}

	metadata, err := parseMetadata(lines)
	if err != nil {
		t.Fatalf("parseMetadata returned an error: %v", err)
	}
	if !reflect.DeepEqual(metadata, expected) {
		t.Errorf("parseMetadata = %v, want %v", metadata, expected)
	}
}

func TestParseTasks(t *testing.T) {
	rawLines := `
## Task
- TODO task1
	- TODO sub task2
		ID: xxx-xxx-xxx
		CREATED: 2023-10-10
	- DONE sub task3 
		ID: yyy-yyy-yyy
		CREATED: 2023-10-10
		CLOSED: 2023-11-10
`
	lines := strings.Split(rawLines, "\n")

	expected := []model.Task{
		{
			Description: "task1",
			Status:      "TODO",
			SubTasks: []model.Task{
				{
					Description: "sub task2",
					Status:      "TODO",
					SubTasks:    []model.Task{},
					Details: map[string]string{
						"ID":      "xxx-xxx-xxx",
						"CREATED": "2023-10-10",
					},
				},
				{
					Description: "sub task3",
					Status:      "DONE",
					SubTasks:    []model.Task{},
					Details: map[string]string{
						"ID":      "yyy-yyy-yyy",
						"CREATED": "2023-10-10",
						"CLOSED":  "2023-11-10",
					},
				},
			},
			Details: map[string]string{},
		},
	}

	tasks, err := parseTasks(lines)
	if err != nil {
		t.Errorf("ParseTasks returned an error: %v", err)
	}

	if diff := cmp.Diff(tasks, expected); diff != "" {
		t.Errorf("%s", diff)

	}
}
