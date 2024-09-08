package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Store[T any](ctx context.Context, coll *mongo.Collection, entity *T) (*T, error) {
	_, err := coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func Find[T any](ctx context.Context, coll *mongo.Collection) ([]*T, error) {
	return []*T{}, nil
}

func FindID[T any](ctx context.Context, coll *mongo.Collection, id primitive.ObjectID) (*T, error) {
	return nil, nil
}
