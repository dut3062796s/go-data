// Package data provides a framework for data access
package data

import (
	"github.com/micro/go-data/model"
)

// Database represents a crud interface
type Database interface {
	Close() error
	Init(...Option) error
	Options() Options
	Model() model.CRUD
}

// Dataflow represents a stream interface
type Dataflow interface {
	Close() error
	Init(...Option) error
	Options() Options
	Model() model.Stream
}

type Option func(*Options)
