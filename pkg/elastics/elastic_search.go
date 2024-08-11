package elastics

import (
	"context"
	"github.com/RandySteven/neo-postman/documentaries"
	documentaries_interfaces "github.com/RandySteven/neo-postman/interfaces/documentaries"
	"github.com/RandySteven/neo-postman/pkg/config"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
	"time"
)

type ESClient struct {
	TestDataDocumentary documentaries_interfaces.TestDataDocumentary
	client              *elasticsearch.Client
}

func NewESClient(config *config.Config) (*ESClient, error) {
	timeout := time.Duration(config.Elasticsearch.Timeout) * time.Second

	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
	}

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://" + config.Elasticsearch.Host + ":" + config.Elasticsearch.Port,
		},
		Username:  config.Elasticsearch.Username,
		Password:  config.Elasticsearch.Password,
		Transport: transport,
	})
	if err != nil {
		return nil, err
	}

	return &ESClient{
		TestDataDocumentary: documentaries.NewTestDataDocumentary(client),
		client:              client,
	}, nil
}

func (e *ESClient) Ping(ctx context.Context) error {
	_, err := e.client.Ping(
		e.client.Ping.WithContext(ctx),
	)
	if err != nil {
		return err
	}
	log.Println("PONG")
	return nil
}
