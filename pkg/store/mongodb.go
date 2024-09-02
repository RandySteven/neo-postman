package store

import (
	"context"
	"fmt"
	"github.com/RandySteven/neo-postman/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client *mongo.Client
}

func NewMongoStore(config *config.Config) (*Store, error) {
	uri := fmt.Sprintf("")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return &Store{
		client: client,
	}, nil
}
