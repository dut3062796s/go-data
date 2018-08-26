// Package mock mocks the Database
package mock

import (
	"errors"
	"sync"

	"github.com/micro/go-data/model"
)

type mockModel struct {
	opts model.Options

	sync.RWMutex
	database map[string]model.Record
}

func (m *mockModel) Init(opts ...model.Option) error {
	for _, o := range opts {
		o(&m.opts)
	}
	return nil
}

func (m *mockModel) Options() model.Options {
	return m.opts
}

func (m *mockModel) Model() model.Model {
	return m
}

func (m *mockModel) String() string {
	return "mock"
}

func (m *mockModel) Read(id string) (model.Record, error) {
	m.RLock()
	defer m.RUnlock()
	r, ok := m.database[id]
	if !ok {
		return nil, errors.New("record not found")
	}
	return r, nil
}

func (m *mockModel) Create(r model.Record) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.database[r.Id()]; ok {
		return errors.New("record already exists")
	}
	m.database[r.Id()] = r
	return nil
}

func (m *mockModel) Update(r model.Record) error {
	m.Lock()
	defer m.Unlock()
	m.database[r.Id()] = r
	return nil
}

func (m *mockModel) Delete(id string) error {
	m.Lock()
	defer m.Unlock()
	delete(m.database, id)
	return nil
}
