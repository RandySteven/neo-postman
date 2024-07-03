package repositories_interfaces

import "context"

type Repositories[T any] interface {
	Save(ctx context.Context, request *T) (result *T, err error)
	FindAll(ctx context.Context) (result []*T, err error)
	FindByID(ctx context.Context, id *uint64) (result *T, err error)
	Update(ctx context.Context, request *T) (result *T, err error)
	Delete(ctx context.Context, id *uint64) (err error)
}
