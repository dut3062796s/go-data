// Package mock mocks the Database
package mock

import (
	"github.com/micro/go-data"
	"github.com/micro/go-data/model"
)

type mockDatabase struct {
	closed chan bool
	opts   data.Options

	model *mockModel
}

type mockRecord struct {
	id   string
	data interface{}
}

func (m *mockDatabase) Close() error {
	select {
	case <-m.closed:
		return nil
	default:
		close(m.closed)
	}
	return nil
}

func (m *mockDatabase) Init(opts ...data.Option) error {
	for _, o := range opts {
		o(&m.opts)
	}
	return nil
}

func (m *mockDatabase) Options() data.Options {
	return m.opts
}

func (m *mockDatabase) Model() model.Model {
	return m.model
}

func (m *mockDatabase) String() string {
	return "mock"
}

// NewRecord creates a new record
func NewRecord(id string, md model.Metadata, data interface{}) model.Record {
	return newRecord(id, md, data)
}

// NewDatabase returns a new mock Database
func NewDatabase(opts ...data.Option) data.Database {
	var options data.Options
	for _, o := range opts {
		o(&options)
	}

	return &mockDatabase{
		opts:   options,
		closed: make(chan bool),
		model: &mockModel{
			database: make(map[string]model.Record),
		},
	}
}
