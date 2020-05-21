package domain

import (
	"fmt"
)

// Task is for storing tasks. Every task should have a title, with
// an optional description, as well as a due date and status.
// Also,  a priority
type Task struct {
	Rowid       int64    `json:"rowid"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	DueDate     Time     `json:"dueDate" db:"dueDate"`
	Status      Status   `json:"status"`
	Priority    Priority `json:"priority"`
	Effort      Duration `json:"effort"`
	Created     Time     `json:"created"`
}

func (t Task) String() string {
	return fmt.Sprintf(`Task:{"title":"%s", dueDate:"%s"}`, t.Title, t.DueDate)
}

// Status is an enum - NotStarted, Doing, Finished
type Status string

// Priority - 5 is highes, 1 is lowest
type Priority uint8
