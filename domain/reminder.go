package domain

// Reminder - A reminder can be created for tasks that are to be reviewed periodically.
// For example, consider you learned a topic today. Now, you want to be reminded of
// it a week later, so as to reinforce learning.
// For that, simply create a reminder for it.
type Reminder struct {
	Rowid        int64  `json:"rowid"`
	Task         Task   `json:"task"`
	ReminderTime Time   `json:"reminderTime" db:"reminderTime"`
	Processed    bool   `json:"processed"`
	Notes        string `json:"notes"`
}
