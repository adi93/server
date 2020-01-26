package api

import (
	"fmt"
	"time"

	"server/domain"
)

// Request section

const (
	titleMaxLength       = 30
	descriptionMaxLength = 600
)

// CreateTaskRequest used for creating a task
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    uint8  `json:"priority"`
	Effort      string `json:"effort"`
}

func (c *CreateTaskRequest) String() string {
	return fmt.Sprintf(`{"title":"%s", "description":"%s", "dueDate":"%s", "Priority": "%d", "Effort": "%s"}`, c.Title, c.Description, c.DueDate, c.Priority, c.Effort)
}

// Validate is for conforming to api.Request interface.
// Title, Priority and DueDate are compulsory columns of a task.
// DueDate should be entered in domain.DateFormat layout, and has to be in future.
func (c *CreateTaskRequest) Validate() error {
	// Validate title
	if c.Title == "" {
		return fmt.Errorf("Cannot have empty title")
	}
	if len(c.Title) > titleMaxLength {
		return fmt.Errorf("Title length cannot be greater than %d", titleMaxLength)
	}

	// Validate priority
	if c.Priority == 0 {
		return fmt.Errorf("Cannot have empty Priority")
	}
	if c.Priority > 5 {
		return fmt.Errorf("Priority ranges from 1 to 5 (1 being the least, 5 being the max)")
	}

	// Validate due date
	if c.DueDate == "" {
		return fmt.Errorf("Cannot have empty due date")
	}
	t, err := time.Parse(domain.DateFormat, c.DueDate)
	if err != nil {
		return err
	}
	if !t.After(time.Now()) {
		return fmt.Errorf("DueDate has to be in future")
	}

	// Vaidate Effort
	if c.Effort != "" {
		if _, err := time.ParseDuration(c.Effort); err != nil {
			return fmt.Errorf("Invalid duration %s", c.Effort)
		}
	} else {
		c.Effort = "24h"
	}

	return nil
}

var _ Request = &CreateTaskRequest{}

// UpdateTaskRequest is for updating a task. It contains the same field as
// CreateTaskRequest, with addition of a status field.
type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    uint8  `json:"priority"`
	Effort      string `json:"effort"`
	Status      string
}

var _ Request = &UpdateTaskRequest{}

func (u *UpdateTaskRequest) String() string {
	return fmt.Sprintf(`{"title":"%s", "description":"%s", "dueDate":"%s", "Priority": "%d", "Effort": "%s", "Status":"%s"}`, u.Title, u.Description, u.DueDate, u.Priority, u.Effort, u.Status)
}

// Validate is for conforming to api.Request interface.
// Status should be a valid domain.Status
func (u *UpdateTaskRequest) Validate() error {
	// Validate title
	if u.Title == "" {
		return fmt.Errorf("Cannot have empty title")
	}
	if len(u.Title) > titleMaxLength {
		return fmt.Errorf("Title length cannot be greater than %d", titleMaxLength)
	}

	// Validate priority
	if u.Priority > 5 || u.Priority < 0 {
		return fmt.Errorf("Priority ranges from 1 to 5 (1 being the least, 5 being the max)")
	}

	// Validate due date
	if u.DueDate == "" {
		t, err := time.Parse(domain.DateFormat, u.DueDate)
		if err != nil {
			return err
		}
		if !t.After(time.Now()) {
			return fmt.Errorf("DueDate has to be in future")
		}
	}

	// Vaidate Effort
	if u.Effort != "" {
		if _, err := time.ParseDuration(u.Effort); err != nil {
			return fmt.Errorf("Invalid duration %s", u.Effort)
		}
	}

	// Status checks
	tempStatus := domain.Status(u.Status)
	if u.Status != "" && !(tempStatus == domain.InProgress || tempStatus == domain.Done || tempStatus == domain.Pending) {
		return fmt.Errorf("Only valid status are: Pending, In-Progress and Done")
	}
	return nil
}

// Response section

// CreateTaskResponse encapsulates taskId and Response
type CreateTaskResponse struct {
	Response `json:"response"`
	TaskID   int64 `json:"task_id"`
}

func (r CreateTaskResponse) String() string {
	return fmt.Sprintf(`{"response": %v, "taskId":%d}`, r.Response.String(), r.TaskID)
}

// GetTaskResponse ...
type GetTaskResponse struct {
	Response `json:"response"`
	Task     domain.Task `json:"task"`
}

func (r GetTaskResponse) String() string {
	return fmt.Sprintf(`{"response": %v, "task":%v}`, r.Response.String(), r.Task)
}

// GetBulkTasksResponse used for returning multiple tasks
type GetBulkTasksResponse struct {
	Response `json:"response"`
	Tasks    []domain.Task `json:"tasks"`
}

func (r GetBulkTasksResponse) String() string {
	return fmt.Sprintf(`{"response": %v, "tasks":%v}`, r.Response.String(), "")
}
