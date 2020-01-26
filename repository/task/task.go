package task

import (
	"context"
	"errors"
	"server/db"
	"server/domain"
	"sync"
)

var (
	taskMu              sync.Mutex
	taskRepoInitialized = false
	taskOnce            sync.Once
	taskRepository      ITaskRepo
)

// Repository is the  accessor for ITaskRepo.
func Repository() ITaskRepo {
	return taskRepository
}

// ITaskRepo implements CRUD operation for Task
type ITaskRepo interface {
	GetTaskByTitle(ctx context.Context, title string) (domain.Task, error)
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	AddTask(ctx context.Context, task domain.Task) (int64, error)
	DeleteTask(ctx context.Context, id int64) error
	DeleteTaskByTitle(ctx context.Context, title string) error
	UpdateTask(ctx context.Context, task domain.Task) error
}

// InitializeTaskRepo ensures that a task repository is created only once
func InitializeTaskRepo(pr ITaskRepo) error {
	taskMu.Lock()
	defer taskMu.Unlock()
	if taskRepoInitialized {
		return errors.New("Initializing task repo again")
	}

	taskOnce.Do(func() {
		taskRepository = pr
		taskRepoInitialized = true
	})
	return nil
}

func InitTaskRepo(handler db.Handler) {
	switch handler.Type() {
	case db.SQLITE:
		InitializeSqlite3TaskRepo(handler)
	default:
		panic("No handler for this type exists")
	}
}
