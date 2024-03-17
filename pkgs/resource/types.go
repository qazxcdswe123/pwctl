package resource

import (
	"context"
)

type ListOptions struct {
	Hostname string
}

type GetOptions struct {
	Name string
	ListOptions
}

type Operations interface {
	Get(ctx context.Context)

	List(ctx context.Context)
}
