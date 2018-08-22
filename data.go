// Package data provides a framework for data access
package data

import (
	"github.com/micro/go-data/model"
)

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
	model.CRUD
	// include Search
	Search(...SearchOption) ([]model.Record, error)
}

type Option func(*Options)

type SearchOption func(*SearchOptions)
