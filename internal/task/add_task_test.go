package task

import (
	"math/rand"
	"testing"
	"time"
	"todotxtplus/internal/model"

	"github.com/google/go-cmp/cmp"
	"github.com/oklog/ulid/v2"
)

func fixedULID() ulid.ULID {
	// 固定されたタイムスタンプとシード値
	fixedTime := time.Unix(1000000, 0)
	fixedSeed := int64(12345)
	// 予測可能な乱数生成器
	entropy := rand.New(rand.NewSource(fixedSeed))
	return ulid.MustNew(ulid.Timestamp(fixedTime), entropy)
}

func TestAddTask(t *testing.T) {
	descriptionStr := "desc message"
	expected := model.TaskList{
		model.Task{
			ID:          fixedULID(),
			Status:      "TODO",
			Description: descriptionStr,
		},
	}

	emptyTasklist := model.TaskList{}

	actual, err := emptyTasklist.AddTask(fixedULID(), descriptionStr)

	if err != nil {
		t.Errorf("addTask returned an error: %v", err)
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("%s", diff)
	}
}

func TestAddTaskToExistingTasks(t *testing.T) {
	expected := model.TaskList{
		model.Task{
			ID:          fixedULID(),
			Status:      "TODO",
			Description: "aaaaaaaaa",
		},
		model.Task{
			ID:          fixedULID(),
			Status:      "TODO",
			Description: "bbbbbbbb",
		},
	}

	list := model.TaskList{
		model.Task{
			ID:          fixedULID(),
			Status:      "TODO",
			Description: "aaaaaaaaa",
		},
	}
	actual, err := list.AddTask(fixedULID(), "bbbbbbbb")

	if err != nil {
		t.Errorf("addTask returned an error: %v", err)
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("%s", diff)
	}
}
