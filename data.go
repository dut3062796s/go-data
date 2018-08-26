// Package data provides a framework for data access
package data

import (
	"github.com/micro/go-data/model"
)

// DB represents a crud interface
type DB interface {
	Close() error
	Init(...Option) error
	Options() Options
	Model() model.Model
}

type Option func(*Options)
