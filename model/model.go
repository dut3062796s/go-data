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

// CRUD is the data model for crud access
type CRUD interface {
	Model
	Read(string) (Record, error)
	Create(Record) error
	Update(Record) error
	Delete(string) error
}

// Stream is the data model for stream processing
type Stream interface {
	Model
	// TODO: Stream interface
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
