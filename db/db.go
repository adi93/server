package db

// Result interface is used to get query execution info
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Row is standard db row
type Row interface {
	Scan(dest ...interface{}) error
	StructScan(dest interface{}) error
}

// Rows are just a collection of Row
type Rows interface {
	Row
	Next() bool
}

// Handler interface to db. Underlying implementation can be sqlite, postgres, etc.
// Should be used in repositories.
type Handler interface {
	Type() Type
	Execute(statement string, args ...interface{}) (Result, error)
	QueryRow(statement string, args ...interface{}) Row
	Query(statement string, args ...interface{}) (Rows, error)
}

// Type is enum for which type of database
type Type string

const (
	// SQLITE represents SQLITE database type
	SQLITE Type = "SQLITE"
)
