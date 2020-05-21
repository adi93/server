package service

import (
	"server/api"
	"server/errors"
	"server/repository/task"

	"context"
	"testing"
)

var taskService ITaskService

func createTaskService() {
	// create a mockk repo
	task.InitializeInMemoryTaskRepo()
	inMemmoryTaskRepo := task.Repository()

	// initialize service
	InitializeTaskService(inMemmoryTaskRepo)
	taskService = TaskService
}

func TestInitializeTaskService(t *testing.T) {
	createTaskService()
	if TaskService == nil {
		t.Fatalf("Could not initialize task service")
	}

}

func TestCreateTask(t *testing.T) {
	tests := []struct {
		req  api.CreateTaskRequest
		resp api.CreateTaskResponse
	}{
		{
			api.CreateTaskRequest{
				Title:       "First",
				Description: "First Task",
				DueDate:     "2006-01-02 03:04:05",
				Priority:    1,
				Effort:      "1.5h",
			},
			api.CreateTaskResponse{
				Response: api.NewStdResponse(),
				TaskID:   0,
			},
		},
		{
			api.CreateTaskRequest{
				Title:       "Second",
				Description: "First Task",
				DueDate:     "2006-01-02 03:04:05",
				Priority:    1,
				Effort:      "1.5h",
			},
			api.CreateTaskResponse{
				Response: api.NewStdResponse(),
				TaskID:   1,
			},
		},
		{
			api.CreateTaskRequest{
				Title:       "Second",
				Description: "First Task",
				DueDate:     "2006-01-02 03:04:05",
				Priority:    1,
				Effort:      "1.5h",
			},
			api.CreateTaskResponse{
				Response: api.NewErrorResponse(errors.ErrorObjectAlreadyExists),
				TaskID:   -1,
			},
		},
	}
	for i, test := range tests {
		actualRespo := taskService.CreateTask(context.Background(), test.req)
		if !actualRespo.Equals(test.resp) {
			t.Logf("Test %d failed. Request: %v, Expected Response: %v, Actual response: %v", i, test.req, test.resp, actualRespo)
			t.Fail()
		}
	}
}
