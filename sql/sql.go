// Package sql provides a SQL interface
package sql

import "database/sql"

// DB is the minimal database connection functionality required. Implemented
// by *sql.DB.
type DB interface {
	Close() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Ping() error
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Transaction defines the TCL operations on a database.
type Transaction interface {
	Commit() error
	Rollback() error
}
