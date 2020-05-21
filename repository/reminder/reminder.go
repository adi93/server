package reminder

import (
	"context"
	"errors"
	"server/db"
	"server/domain"
	"sync"
)

var (
	reminderMu              sync.Mutex
	reminderRepoInitialized = false
	reminderOnce            sync.Once
	reminderRepository      IReminderRepo
)

// Repository is the  accessor for IReminderRepo.
func Repository() IReminderRepo {
	return reminderRepository
}

// IReminderRepo implements CRUD operation for Reminder
type IReminderRepo interface {
	GetReminderForTask(ctx context.Context, taskId int64) (domain.Reminder, error)
	AddReminder(ctx context.Context, reminder domain.Reminder) (int64, error)
	DeleteReminder(ctx context.Context, id int64) error
	UpdateReminder(ctx context.Context, reminder domain.Reminder) error
}

// InitializeReminderRepo ensures that a reminder repository is created only once
func InitializeReminderRepo(pr IReminderRepo) error {
	reminderMu.Lock()
	defer reminderMu.Unlock()
	if reminderRepoInitialized {
		return errors.New("Initializing reminder repo again")
	}

	reminderOnce.Do(func() {
		reminderRepository = pr
		reminderRepoInitialized = true
	})
	return nil
}

func InitReminderRepo(handler db.Handler) {
	switch handler.Type() {
	case db.SQLITE:
		InitializeSqlite3ReminderRepo(handler)
	default:
		panic("No handler for this type exists")
	}
}
