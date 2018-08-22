package data

import (
	"context"

	"github.com/micro/go-data/model"
)

type Options struct {
	Database string
	Table    string

	// For alternative options
	Context context.Context
}

type SearchOptions struct {
	Metadata model.Metadata
	Limit    int64
	Offset   int64
}
