package caches_interfaces

import "context"

type Cache[T any] interface {
	Set(ctx context.Context, key string, value *T) (err error)
	Get(ctx context.Context, key string) (value *T, ok bool)
	Del(ctx context.Context, key string) (err error)
}
