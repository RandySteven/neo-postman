package elastics

import (
	"github.com/RandySteven/neo-postman/documentaries"
	documentaries_interfaces "github.com/RandySteven/neo-postman/interfaces/documentaries"
	"github.com/RandySteven/neo-postman/pkg/config"
	"github.com/elastic/go-elasticsearch/v8"
)

type ESClient struct {
	TestDataDocumentary documentaries_interfaces.TestDataDocumentary
	client              *elasticsearch.Client
}

func NewESClient(config *config.Config) (*ESClient, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			config.Elasticsearch.Host + ":" + config.Elasticsearch.Port,
		},
		Username: config.Elasticsearch.Username,
		Password: config.Elasticsearch.Password,
	})
	if err != nil {
		return nil, err
	}
	return &ESClient{
		TestDataDocumentary: documentaries.NewTestDataDocumentary(client),
		client:              client,
	}, nil
}
