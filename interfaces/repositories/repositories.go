package repositories_interfaces

import (
	"context"
	"database/sql"
)

type (
	Repositories[T any] interface {
		Save(ctx context.Context, request *T) (result *T, err error)
		FindAll(ctx context.Context) (result []*T, err error)
		FindByID(ctx context.Context, id uint64) (result *T, err error)
		Update(ctx context.Context, request *T) (result *T, err error)
		Delete(ctx context.Context, id uint64) (err error)
	}

	Transaction[T any] interface {
		Begin(ctx context.Context) (err error)
		Commit(ctx context.Context) (err error)
		Rollback(ctx context.Context) (err error)
		SetTx(tx *sql.Tx)
	}
)
