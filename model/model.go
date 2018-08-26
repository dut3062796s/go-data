// Package model represents the data model
package model

// Model is the basis for a data model
type Model interface {
	// Initialise options
	Init(...Option) error
	// Get Options
	Options() Options
	// Name of the model
	String() string
}

// crud represents a data model
type CRUD interface {
	Model
	Read(string) (Record, error)
	Create(Record) error
	Update(Record) error
	Delete(string) error
}

type Metadata map[string]interface{}

type Options struct{}

type Option func(*Options)

type Record interface {
	Id() string
	Created() int64
	Updated() int64
	Metadata() Metadata
	Bytes() []byte
	Scan(v interface{}) error
}
