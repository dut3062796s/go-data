// Package mock mocks the Database
package mock

import (
	"errors"
	"sync"

	"github.com/micro/go-data"
	"github.com/micro/go-data/model"
)

type mockDatabase struct {
	closed chan bool
	opts   data.Options

	sync.RWMutex
	database map[string]model.Record
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
	return m
}

func (m *mockDatabase) String() string {
	return "mock"
}

func (m *mockDatabase) Read(id string) (model.Record, error) {
	m.RLock()
	defer m.RUnlock()
	r, ok := m.database[id]
	if !ok {
		return nil, errors.New("record not found")
	}
	return r, nil
}

func (m *mockDatabase) Create(r model.Record) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.database[r.Id()]; ok {
		return errors.New("record already exists")
	}
	m.database[r.Id()] = r
	return nil
}

func (m *mockDatabase) Update(r model.Record) error {
	m.Lock()
	defer m.Unlock()
	m.database[r.Id()] = r
	return nil
}

func (m *mockDatabase) Delete(id string) error {
	m.Lock()
	defer m.Unlock()
	delete(m.database, id)
	return nil
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
		opts:     options,
		database: make(map[string]model.Record),
		closed:   make(chan bool),
	}
}
