// Package blob represents blob storage
package blob

import (
	"io"
)

// Store represents blob storage
type Store interface {
	// Operations on Buckets
	Create(string) error
	Delete(string) error
	Get(string) (Bucket, error)
	List() ([]Bucket, error)
}

// type Bucket represents a bucket in the store
type Bucket interface {
	// Operations on Objects
	Create(string) (Blob, error)
	Delete(string) error
	Get(string) (Blob, error)
}

type Blob interface {
	Name() string
	Metadata() map[string]interface{}
	Read() io.ReadCloser
	Write() io.Writer
}
