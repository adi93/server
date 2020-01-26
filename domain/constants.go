package domain

const (
	// DateFormat is the standard parsing format for all kinds of dates.
	DateFormat = "2006-01-02 03:04:05"
)

const (
	// Pending means task has not been picked up
	Pending Status = "Pending"
	// InProgress means task is active and being worked upon
	InProgress Status = "In-Progress"
	// Done means taks is done
	Done Status = "Done"
)
