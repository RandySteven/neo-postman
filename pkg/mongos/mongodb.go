package mongos

import (
	"context"
	"fmt"
	"github.com/RandySteven/neo-postman/collections"
	collections_interfaces "github.com/RandySteven/neo-postman/interfaces/collections"
	"github.com/RandySteven/neo-postman/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	ApiContentDetail collections_interfaces.ApiContentDetailCollection
	client           *mongo.Client
}

func NewMongoStore(config *config.Config) (*MongoDB, error) {
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

	database := client.Database("")

	return &MongoDB{
		ApiContentDetail: collections.NewApiContentDetailCollection(database),
		client:           client,
	}, nil
}

func (m *MongoDB) Ping() error {
	if err := m.client.Ping(context.TODO(), nil); err != nil {
		return err
	}
	return nil
}
