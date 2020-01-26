package service

import (
	"context"
	"log"
	"server/errors"
	"time"

	"server/api"
	"server/domain"
	"server/repository/task"
)

var (
	// TaskService is the accessor of ITaskService.
	// Initialize it with InitializeTaskService
	TaskService     ITaskService
	taskServiceCode = "TaskService"
)

// init just adds a taskServiceBuilder to Initializers map
func init() {
	log.Printf("Initializing task service")
	builder := NewBaseBuilder(taskServiceCode, false)
	b := taskServiceBuilder{&builder}
	Initializers[taskServiceCode] = &b
	log.Printf("Initialized task service")
}

// ITaskService is the interface for CRUD operations of tasks
type ITaskService interface {
	CreateTask(ctx context.Context, r api.CreateTaskRequest) api.CreateTaskResponse
	GetTask(ctx context.Context, name string) api.GetTaskResponse
	GetAllTasks(ctx context.Context) api.GetBulkTasksResponse
	DeleteTask(ctx context.Context, name string) api.Response
	UpdateTask(ctx context.Context, r api.UpdateTaskRequest) api.Response
}

// InitializeTaskService initializes the task service
func InitializeTaskService(repo task.ITaskRepo) error {
	builder := Initializers[taskServiceCode]
	if err := build(builder, repo); err != nil {
		return err
	}
	return nil
}

type taskServiceBuilder struct {
	*BaseBuilder
}

// Build is used to initialize channel service
func (tsb *taskServiceBuilder) Build(args ...interface{}) error {
	if len(args) != 1 {
		return errors.ErrorArgumentMismatch
	}
	value := args[0]
	repo, ok := value.(task.ITaskRepo)
	if !ok {
		return errors.ErrorInvalidType
	}
	TaskService = TaskServiceImpl{repo}
	return nil
}

// TaskServiceImpl implements ITaskService
type TaskServiceImpl struct {
	repo task.ITaskRepo
}

// CreateTask creates task and stores in the repository
func (ts TaskServiceImpl) CreateTask(ctx context.Context, r api.CreateTaskRequest) api.CreateTaskResponse {
	dueDate, _ := time.Parse(domain.DateFormat, r.DueDate)
	effort, _ := time.ParseDuration(r.Effort)
	task := domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Priority:    domain.Priority(r.Priority),
		DueDate:     domain.Time(dueDate),
		Effort:      domain.Duration(effort),
		Status:      domain.Pending,
	}
	id, err := ts.repo.AddTask(ctx, task)
	if err != nil {
		return api.CreateTaskResponse{Response: api.NewErrorResponse(err), TaskID: -1}
	}
	return api.CreateTaskResponse{Response: api.NewStdResponse(), TaskID: id}

}

// GetTask gets a task by it's title
func (ts TaskServiceImpl) GetTask(ctx context.Context, title string) api.GetTaskResponse {
	task, err := ts.repo.GetTaskByTitle(ctx, title)
	if err != nil {
		return api.GetTaskResponse{Response: api.NewErrorResponse(err), Task: domain.Task{}}
	}
	log.Printf("Task time: %v, %v", task.DueDate, task.Created)
	return api.GetTaskResponse{Response: api.NewStdResponse(), Task: task}
}

// GetAllTasks gets all the tasks from a repository. Don't use this when db gets too big, use
// GetAllPaginatedTasks, which supports pagination.
func (ts TaskServiceImpl) GetAllTasks(ctx context.Context) api.GetBulkTasksResponse {
	tasks, err := ts.repo.GetAllTasks(ctx)
	if err != nil {
		return api.GetBulkTasksResponse{Response: api.NewErrorResponse(err), Tasks: []domain.Task{}}
	}
	return api.GetBulkTasksResponse{Response: api.NewStdResponse(), Tasks: tasks}
}

// DeleteTask deletes a task. I don't plan to use it much.
func (ts TaskServiceImpl) DeleteTask(ctx context.Context, title string) api.Response {
	task, err := ts.repo.GetTaskByTitle(ctx, title)
	if err != nil {
		return api.NewErrorResponse(err)
	}

	err = ts.repo.DeleteTask(ctx, task.Rowid)
	if err != nil {
		return api.NewErrorResponse(err)
	}
	return api.NewStdResponse()
}

// UpdateTask deletes a task. I don't plan to use it much.
func (ts TaskServiceImpl) UpdateTask(ctx context.Context, r api.UpdateTaskRequest) api.Response {
	task, err := ts.repo.GetTaskByTitle(ctx, r.Title)
	if err != nil {
		return api.NewErrorResponse(err)
	}

	// set new values
	if r.Status != "" {
		task.Status = domain.Status(r.Status)
	}
	if r.Description != "" {
		task.Description = r.Description
	}
	if r.DueDate != "" {
		dueDate, _ := time.Parse(domain.DateFormat, r.DueDate)
		task.DueDate = domain.Time(dueDate)
	}
	if r.Priority != 0 {
		task.Priority = domain.Priority(r.Priority)
	}
	if r.Effort != "" {
		effort, _ := time.ParseDuration(r.Effort)
		task.Effort = domain.Duration(effort)
	}
	if r.Status != "" {
		task.Status = domain.Status(r.Status)
	}

	err = ts.repo.UpdateTask(ctx, task)
	if err != nil {
		return api.NewErrorResponse(err)
	}
	return api.NewStdResponse()
}
