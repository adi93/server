package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // For sqlite3 driver
)

// SqliteRow - implements Row interface
type SqliteRow struct {
	Row *sqlx.Row
}

// Scan is a wrapper around sql.Scan. Used for scanning
// db values in a row.
func (r SqliteRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

// StructScan is a wrapper around sqlx.Row.StructScan
func (r SqliteRow) StructScan(dest interface{}) error {
	return r.Row.StructScan(dest)
}

// SqliteRows - implements Rows iterface
type SqliteRows struct {
	Rows *sqlx.Rows
}

// Scan is a wrapper around sql.Scan. Used for scanning
// db values in a row.
func (r SqliteRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// StructScan is a wrapper around sqlx.Rows.StructScan
func (r SqliteRows) StructScan(dest interface{}) error {
	return r.Rows.StructScan(dest)
}

// Next is a wrapper around sql.Next. Used for iterating
// on multiple rows
func (r SqliteRows) Next() bool {
	return r.Rows.Next()
}

// SqliteHandler implements Handler intrface.
type SqliteHandler struct {
	Conn *sqlx.DB
}

// Type of SqliteHandler is SQLITE
func (handler *SqliteHandler) Type() Type {
	return SQLITE
}

// Execute is for non-select queries.
func (handler *SqliteHandler) Execute(statement string, args ...interface{}) (Result, error) {
	stmt, err := handler.Conn.Preparex(statement)
	if err != nil {
		return nil, err
	}
	// return stmt.MustExec(args...), nil
	return stmt.Exec(args...)
}

// Query is for those queries which return only 1 row.
func (handler *SqliteHandler) Query(statement string, args ...interface{}) (Rows, error) {
	r, err := handler.Conn.Queryx(statement, args...)
	if err != nil {
		return new(SqliteRows), err
	}
	rows := new(SqliteRows)
	rows.Rows = r
	return rows, nil
}

// QueryRow is for those select queries which return multiple rows
func (handler *SqliteHandler) QueryRow(statement string, args ...interface{}) Row {
	r := handler.Conn.QueryRowx(statement, args...)
	row := new(SqliteRow)
	row.Row = r
	return row
}

// NewSqliteHandler returns an SqliteHandler
func NewSqliteHandler(dbfileName string) *SqliteHandler {
	conn, _ := sqlx.Open("sqlite3", dbfileName)
	sqliteHandler := new(SqliteHandler)
	sqliteHandler.Conn = conn
	return sqliteHandler
}

// InMemorySqliteHandler is an in memory db, useful for testing
var InMemorySqliteHandler = NewSqliteHandler(":memory:")
