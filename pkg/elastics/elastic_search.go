package elastics

import (
	"github.com/RandySteven/neo-postman/pkg/config"
	"github.com/elastic/go-elasticsearch/v8"
)

type ESClient struct {
	client *elasticsearch.Client
}

func NewESClient(config *config.Config) (*ESClient, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{})
	if err != nil {
		return nil, err
	}
	return &ESClient{
		client: client,
	}, nil
}
