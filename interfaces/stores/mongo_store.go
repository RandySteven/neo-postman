package stores

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoStore[T any] interface {
	Store(ctx context.Context, entity *T) (result *T, err error)
	FindAll(ctx context.Context) (result []*T, err error)
	FindById(ctx context.Context, id primitive.ObjectID) (result *T, err error)
	DeleteById(ctx context.Context, id primitive.ObjectID) (err error)
	UpdateById(ctx context.Context, entity *T) (result *T, err error)
}
