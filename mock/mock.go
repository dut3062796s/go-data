// Package mock mocks the DB
package mock

import (
	"errors"
	"sync"

	"github.com/micro/go-data"
)

type mockDB struct {
	closed chan bool
	opts   data.Options

	sync.RWMutex
	database map[string]data.Record
}

type mockRecord struct {
	id   string
	data interface{}
}

func (m *mockDB) Close() error {
	select {
	case <-m.closed:
		return nil
	default:
		close(m.closed)
	}
	return nil
}

func (m *mockDB) Init(opts ...data.Option) error {
	for _, o := range opts {
		o(&m.opts)
	}
	return nil
}

func (m *mockDB) Options() data.Options {
	return m.opts
}

func (m *mockDB) String() string {
	return "mock"
}

func (m *mockDB) Read(id string) (data.Record, error) {
	m.RLock()
	defer m.RUnlock()
	r, ok := m.database[id]
	if !ok {
		return nil, errors.New("record not found")
	}
	return r, nil
}

func (m *mockDB) Create(r data.Record) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.database[r.Id()]; ok {
		return errors.New("record already exists")
	}
	m.database[r.Id()] = r
	return nil
}

func (m *mockDB) Update(r data.Record) error {
	m.Lock()
	defer m.Unlock()
	m.database[r.Id()] = r
	return nil
}

func (m *mockDB) Delete(id string) error {
	m.Lock()
	defer m.Unlock()
	delete(m.database, id)
	return nil
}

func (m *mockDB) Search(opts ...data.SearchOption) ([]data.Record, error) {
	return nil, errors.New("not implemented")
}

// NewRecord creates a new record
func NewRecord(id string, md data.Metadata, data interface{}) data.Record {
	return newRecord(id, md, data)
}

// NewDB returns a new mock DB
func NewDB(opts ...data.Option) data.DB {
	var options data.Options
	for _, o := range opts {
		o(&options)
	}

	return &mockDB{
		opts:     options,
		database: make(map[string]data.Record),
		closed:   make(chan bool),
	}
}
