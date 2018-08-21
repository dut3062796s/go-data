// Package sql provides a SQL interface
package sql

type DB interface {
	Exec(query string, args ...interface{}) error
	Query(query string, args ...interface{}) (Rows, error)
	Close() error
}

type Rows interface {
	Close() error
	Next() (Row, error)
}

type Row interface {
	Scan(v interface{}) error
}
