package dbAdapter

import (
	"context"
)

type Options struct {
	Limit    int64
	CursorID interface{}
}

type Table interface {
	QueryOne(ctx context.Context, id string, result interface{}) error
	//QueryPagination(params interface{}, opts Options) ([]Result, error)
}

type Database interface {
	UseTable(tabeName string) Table
}
