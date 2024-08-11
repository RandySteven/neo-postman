package documentaries_interfaces

import "context"

type Documentary[T any] interface {
	MakeAnIndex(ctx context.Context, document *T) (err error)
	GetIndex(ctx context.Context, id uint64) (document *T, err error)
	SearchDocument(ctx context.Context, query string) (t []*T, err error)
	DeletingDocument(ctx context.Context) (err error)
	DeletingIndex(ctx context.Context) (err error)
}
