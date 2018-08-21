// Package data provides a framework for data access
package data

// DB represents a crud interface
type DB interface {
	CRUD
	Close() error
	Init(...Option) error
	Options() Options
	String() string
}

// crud represents a top-level data model
type CRUD interface {
	Read(string) (Record, error)
	Create(Record) error
	Update(Record) error
	Delete(string) error
	Search(...SearchOption) ([]Record, error)
}

type Option func(*Options)

type SearchOption func(*SearchOptions)

type Metadata map[string]interface{}

type Record interface {
	Id() string
	Created() int64
	Updated() int64
	Metadata() Metadata
	Bytes() []byte
	Scan(v interface{}) error
}
