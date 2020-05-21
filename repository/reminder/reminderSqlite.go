package reminder

import (
	"log"
	"server/db"
	"server/domain"

	"context"
	"fmt"
)

// reminderRepositorySqlite implements IReminderRepo interface for sqlite db
type reminderRepositorySqlite struct {
	dbHandler db.Handler
}

// InitializeSqlite3ReminderRepo creates a sqlite db handler, and then calls
// InitializeReminderRepo function, which ensures that only one reminder repository is ever initialized
//
// It doesnt' return anything except an error.
// Sample usage
//
//		import (
//			"server/repository/reminder"
//			reminderImpl "server/repository/reminder/impl"
//		)
//
//		dbHandler := newHandler()
//		reminderImpl.InitializeSqlite3ReminderRepo(dbHandler)
//		reminderRepo := Repository()
//
func InitializeSqlite3ReminderRepo(db db.Handler) error {
	if err := InitializeReminderRepo(newReminderRepoSqlite(db)); err != nil {
		return err
	}
	return nil
}

// newReminderRepoSqlite returns an sqlite3 reminder repository
func newReminderRepoSqlite(db db.Handler) IReminderRepo {
	dbReminderRepo := new(reminderRepositorySqlite)
	dbReminderRepo.dbHandler = db
	return dbReminderRepo
}

var _ IReminderRepo = reminderRepositorySqlite{}

// GetReminderByTitle gets a reminder by its task's title
func (pr reminderRepositorySqlite) GetReminderForTask(ctx context.Context, taskId int64) (reminder domain.Reminder, err error) {
	defer func() {
		if r := recover(); r != nil {
			reminder = domain.Reminder{}
			err = fmt.Errorf("%v", r)
		}
	}()

	row := pr.dbHandler.QueryRow("SELECT * FROM reminder LEFt JOIN task on reminder.taskId = task.rowid WHERE reminder.taskId =? LIMIT 1", taskId)
	err = row.StructScan(&reminder)
	return
}

// AddReminder saves a reminder in db. Returns the Row id of the reminder created, error if no reminder was created
func (pr reminderRepositorySqlite) AddReminder(ctx context.Context, reminder domain.Reminder) (id int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			id = 0
			err = fmt.Errorf("%v", r)
		}
	}()

	res, err := pr.dbHandler.Execute("INSERT INTO reminder (title, description, dueDate, status, priority, effort ) VALUES($1, $2, $3, $4, $5, $6)", reminder.Title, reminder.Description, reminder.DueDate.String(), reminder.Status, reminder.Priority, reminder.Effort)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return
}

// DeleteReminder deletes a reminder by its id
func (pr reminderRepositorySqlite) DeleteReminder(ctx context.Context, id int64) error {
	_, err := pr.dbHandler.Execute("DELETE FROM reminder WHERE rowid = ?", id)
	return err
}

func (pr reminderRepositorySqlite) UpdateReminder(ctx context.Context, reminder domain.Reminder) error {
	_, err := pr.dbHandler.Execute("UPDATE reminder SET description = $1, dueDate = $2, status = $3, priority = $4, effort = $5", reminder.Description, reminder.DueDate.String(), reminder.Status, reminder.Priority, reminder.Effort)
	return err
}
