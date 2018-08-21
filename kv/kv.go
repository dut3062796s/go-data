// Package kv provides a keyvalue interface
package kv

// KV provides a simple keyvalue interface
type KV interface {
	Close() error
	Get(key []byte) (Value, error)
	Put(key []byte, val interface{}) error
	Del(key []byte) error
	Scan(start, end []byte) Scanner
}

// Scanner is used for range scanning
type Scanner interface {
	Close() error
	Next() bool
	Key() []byte
	Val() Value
	Err() error
}

// Value is returned from the DB
type Value interface {
	Key() []byte
	Bytes() []byte
	Scan(v interface{}) error
}
