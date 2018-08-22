// Package model represents the data model
package model

// crud represents a data model
type CRUD interface {
	Read(string) (Record, error)
	Create(Record) error
	Update(Record) error
	Delete(string) error
}

type Metadata map[string]interface{}

type Record interface {
	Id() string
	Created() int64
	Updated() int64
	Metadata() Metadata
	Bytes() []byte
	Scan(v interface{}) error
}
