package data

import (
	"context"
)

type Options struct {
	Database string
	Table    string

	// For alternative options
	Context context.Context
}

type SearchOptions struct {
	Metadata Metadata
	Limit    int64
	Offset   int64
}
