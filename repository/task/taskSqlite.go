package task

import (
	"log"
	"server/db"
	"server/domain"

	"context"
	"fmt"
)

// taskRepositorySqlite implements ITaskRepo interface for sqlite db
type taskRepositorySqlite struct {
	dbHandler db.Handler
}

// InitializeSqlite3TaskRepo creates a sqlite db handler, and then calls
// InitializeTaskRepo function, which ensures that only one task repository is ever initialized
//
// It doesnt' return anything except an error.
// Sample usage
//
//		import (
//			"server/repository/task"
//			taskImpl "server/repository/task/impl"
//		)
//
//		dbHandler := newHandler()
//		taskImpl.InitializeSqlite3TaskRepo(dbHandler)
//		taskRepo := Repository()
//
func InitializeSqlite3TaskRepo(db db.Handler) error {
	if err := InitializeTaskRepo(newTaskRepoSqlite(db)); err != nil {
		return err
	}
	return nil
}

// newTaskRepoSqlite returns an sqlite3 task repository
func newTaskRepoSqlite(db db.Handler) ITaskRepo {
	dbTaskRepo := new(taskRepositorySqlite)
	dbTaskRepo.dbHandler = db
	return dbTaskRepo
}

var _ ITaskRepo = taskRepositorySqlite{}

// GetTaskByTitle gets a task by its title
func (pr taskRepositorySqlite) GetTaskByTitle(ctx context.Context, title string) (task domain.Task, err error) {
	defer func() {
		if r := recover(); r != nil {
			task = domain.Task{}
			err = fmt.Errorf("%v", r)
		}
	}()

	row := pr.dbHandler.QueryRow("SELECT * FROM task WHERE title =?", title)
	err = row.StructScan(&task)
	if err != nil {
		log.Printf("Title: %s", title)
	}
	return
}

// GetAllTasks returns all tasks
func (pr taskRepositorySqlite) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	var tasks []domain.Task
	rows, err := pr.dbHandler.Query("SELECT * FROM task")
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		var p domain.Task
		err := rows.StructScan(&p)
		if err != nil {
			return make([]domain.Task, 0), err
		}

		tasks = append(tasks, p)
	}
	return tasks, nil
}

// AddTask saves a task in db. Returns the Row id of the task created, error if no task was created
func (pr taskRepositorySqlite) AddTask(ctx context.Context, task domain.Task) (id int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			id = 0
			err = fmt.Errorf("%v", r)
		}
	}()

	res, err := pr.dbHandler.Execute("INSERT INTO task (title, description, dueDate, status, priority, effort ) VALUES($1, $2, $3, $4, $5, $6)", task.Title, task.Description, task.DueDate.String(), task.Status, task.Priority, task.Effort)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return
}

// DeleteTask deletes a task by its id
func (pr taskRepositorySqlite) DeleteTask(ctx context.Context, id int64) error {
	_, err := pr.dbHandler.Execute("DELETE FROM task WHERE rowid = ?", id)
	return err
}

// DeleteTaskByTitle deletes a task by its title
func (pr taskRepositorySqlite) DeleteTaskByTitle(ctx context.Context, title string) error {
	_, err := pr.dbHandler.Execute("DELETE FROM task WHERE title = ?", title)
	return err
}

func (pr taskRepositorySqlite) UpdateTask(ctx context.Context, task domain.Task) error {
	_, err := pr.dbHandler.Execute("UPDATE task SET description = $1, dueDate = $2, status = $3, priority = $4, effort = $5", task.Description, task.DueDate.String(), task.Status, task.Priority, task.Effort)
	return err
}
