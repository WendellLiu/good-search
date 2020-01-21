package dbAdapter

import (
	"context"
)

type Options struct {
	Limit    int64
	CursorID string
}

type Table interface {
	QueryOne(ctx context.Context, id string, result interface{}) error
	QueryPagination(ctx context.Context, params map[string]interface{}, opts Options, results interface{}) error
	AllCount(ctx context.Context) (count int64, err error)
}

type Database interface {
	UseTable(tabeName string) Table
}
