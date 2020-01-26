package task

import (
	"server/domain"

	"context"
)

// ContextKey ...
type ContextKey string

const (
	// Debug is used in testing
	Debug ContextKey = "debug"
)

// InitializeMockTaskRepo Passing "debug" in context turns on debugMap mode, where response
// can be controlled by it.
func InitializeMockTaskRepo() {
	repo := mockTaskRepository{}
	InitializeTaskRepo(&repo)
}

type mockTaskRepository struct {
}

// GetTaskByTitle is default
func (pr *mockTaskRepository) GetTaskByTitle(ctx context.Context, title string) (domain.Task, error) {
	if debugMap, ok := ctx.Value(Debug).(map[string]interface{}); ok {
		task, _ := debugMap["task"].(domain.Task)
		err, _ := debugMap["error"].(error)
		return task, err
	}

	return domain.Task{}, nil
}

// GetAllTasks is default
func (pr *mockTaskRepository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	if debugMap, ok := ctx.Value(Debug).(map[string]interface{}); ok {
		tasks, _ := debugMap["tasks"].([]domain.Task)
		err, _ := debugMap["error"].(error)
		return tasks, err
	}

	return make([]domain.Task, 0), nil
}

// AddTask is default
func (pr *mockTaskRepository) AddTask(ctx context.Context, task domain.Task) (id int64, err error) {
	if debugMap, ok := ctx.Value(Debug).(map[string]interface{}); ok {
		id, _ := debugMap["id"].(int64)
		err, _ := debugMap["error"].(error)
		return id, err
	}

	return 0, nil
}

// DeleteTask is default
func (pr *mockTaskRepository) DeleteTask(ctx context.Context, id int64) error {
	if debugMap, ok := ctx.Value(Debug).(map[string]interface{}); ok {
		err, _ := debugMap["error"].(error)
		return err
	}

	return nil
}

// DeleteTaskByTitle is default
func (pr *mockTaskRepository) DeleteTaskByTitle(ctx context.Context, title string) error {
	if debugMap, ok := ctx.Value(Debug).(map[string]interface{}); ok {
		err, _ := debugMap["error"].(error)
		return err
	}
	return nil
}

// UpdateTask adds task image
func (pr *mockTaskRepository) UpdateTask(ctx context.Context, task domain.Task) (err error) {
	if debugMap, ok := ctx.Value(Debug).(map[string]interface{}); ok {
		err, _ := debugMap["error"].(error)
		return err
	}
	return nil

}
