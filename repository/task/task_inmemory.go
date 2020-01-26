package task

import (
	"server/domain"
	"server/errors"

	"context"
)

// InitializeInMemoryTaskRepo can be used for testing.
func InitializeInMemoryTaskRepo() {
	m := make(map[string]domain.Task, 0)
	im := make(map[int64]domain.Task, 0)
	repo := inMemoryTaskRepository{m, im}
	InitializeTaskRepo(&repo)
}

type inMemoryTaskRepository struct {
	m  map[string]domain.Task
	im map[int64]domain.Task
}

// GetTaskByTitle is default
func (pr *inMemoryTaskRepository) GetTaskByTitle(ctx context.Context, title string) (domain.Task, error) {
	p, ok := pr.m[title]
	if !ok {
		return domain.Task{}, errors.ErrorObjectNotFound
	}
	return p, nil
}

// GetAllTasks is default
func (pr *inMemoryTaskRepository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	tasksList := make([]domain.Task, len(pr.m))
	for _, v := range pr.m {
		tasksList = append(tasksList, v)
	}
	return tasksList, nil
}

// AddTask is default
func (pr *inMemoryTaskRepository) AddTask(ctx context.Context, task domain.Task) (id int64, err error) {
	_, ok := pr.m[task.Title]
	if ok {
		return 0, errors.ErrorObjectAlreadyExists
	}
	pr.m[task.Title] = task
	pr.im[task.Rowid] = task
	return task.Rowid, nil
}

// DeleteTask is default
func (pr *inMemoryTaskRepository) DeleteTask(ctx context.Context, id int64) error {
	if _, ok := pr.im[id]; !ok {
		return errors.ErrorObjectNotFound
	}
	task := pr.im[id]
	delete(pr.im, id)
	delete(pr.m, task.Title)
	return nil
}

// DeleteTaskByTitle is default
func (pr *inMemoryTaskRepository) DeleteTaskByTitle(ctx context.Context, title string) error {
	if _, ok := pr.m[title]; !ok {
		return errors.ErrorObjectNotFound
	}
	task := pr.m[title]
	delete(pr.im, task.Rowid)
	delete(pr.m, task.Title)
	return nil
}

// UpdateTask adds task image
func (pr *inMemoryTaskRepository) UpdateTask(ctx context.Context, task domain.Task) (err error) {
	pr.m[task.Title] = task
	pr.im[task.Rowid] = task
	return nil
}
