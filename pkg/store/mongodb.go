package store

import (
	"github.com/RandySteven/neo-postman/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	client *mongo.Client
}

func NewMongoStore(config *config.Config) (*Store, error) {
	return nil, nil
}
