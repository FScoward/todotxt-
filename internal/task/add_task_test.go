package task

import (
	"testing"
	"todotxtplus/internal/model"

	"github.com/google/go-cmp/cmp"
)

func TestAddTask(t *testing.T) {
	descriptionStr := "desc message"
	expected := model.TaskList{
		model.Task{
			Description: descriptionStr,
		},
	}

	emptyTasklist := model.TaskList{}
	actual, err := emptyTasklist.AddTask(descriptionStr)

	if err != nil {
		t.Errorf("addTask returned an error: %v", err)
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("%s", diff)
	}
}
