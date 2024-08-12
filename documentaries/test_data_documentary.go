package documentaries

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/RandySteven/neo-postman/entities/models"
	documentaries_interfaces "github.com/RandySteven/neo-postman/interfaces/documentaries"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

type testDataDocumentary struct {
	client *elasticsearch.Client
}

type MultiMatchQuery struct {
	Query  string   `json:"query"`
	Fields []string `json:"fields"`
}

type Query struct {
	MultiMatch MultiMatchQuery `json:"multi_match"`
}

type SearchRequest struct {
	Query Query `json:"query"`
}

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			Source models.TestData `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (t *testDataDocumentary) createIndexIfNotExists(ctx context.Context, indexName string) error {
	exists, err := t.client.Indices.Exists([]string{indexName})
	if err != nil {
		return fmt.Errorf("error checking index existence: %w", err)
	}
	if exists.StatusCode == 404 {
		_, err = t.client.Indices.Create(indexName)
		if err != nil {
			return fmt.Errorf("error creating index: %w", err)
		}
	}
	return nil
}

func (t *testDataDocumentary) MakeAnIndex(ctx context.Context, document *models.TestData) (err error) {
	if err := t.createIndexIfNotExists(ctx, "test-data-index"); err != nil {
		return err
	}

	data, err := json.Marshal(document)
	if err != nil {
		return fmt.Errorf("error marshaling document: %w", err)
	}

	_, err = t.client.Index("test-data-index", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("error indexing document: %w", err)
	}

	return nil
}

func (t *testDataDocumentary) GetIndex(ctx context.Context, id uint64) (document *models.TestData, err error) {
	res, err := t.client.Get("test-data-index", fmt.Sprintf("%d", id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf(res.String())
	}
	document = &models.TestData{}
	err = json.NewDecoder(res.Body).Decode(&document)
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (t *testDataDocumentary) SearchDocument(ctx context.Context, query string) (result []*models.TestData, err error) {
	// Construct the query using a struct
	searchRequest := SearchRequest{
		Query: Query{
			MultiMatch: MultiMatchQuery{
				Query: query,
				Fields: []string{
					"id", "host", "method", "uri", "description",
					"request_header", "request_body", "actual_response_code", "expected_response_code",
					"result_status",
				},
			},
		},
	}

	body, err := json.Marshal(searchRequest)
	if err != nil {
		log.Println("err : ", err)
		return nil, fmt.Errorf("error marshaling search request: %w", err)
	}
	log.Printf("Search Query: %s\n", string(body))

	// Perform the search
	res, err := t.client.Search(
		t.client.Search.WithContext(ctx),
		t.client.Search.WithIndex("test-data-index"),
		t.client.Search.WithBody(bytes.NewBuffer(body)),
		t.client.Search.WithPretty(),
	)
	if err != nil {
		log.Println("executor search err : ", err)
		return nil, fmt.Errorf("error executing search: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("search request failed: %s", res.String())
	}

	// Decode the search response
	var searchResponse SearchResponse
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		log.Println("err decoder : ", err)
		return nil, fmt.Errorf("error decoding search response: %w", err)
	}

	// Extract the hits into the result slice
	for _, hit := range searchResponse.Hits.Hits {
		result = append(result, &hit.Source)
	}

	return result, nil
}

func (t *testDataDocumentary) DeletingDocument(ctx context.Context, id uint64) error {
	res, err := t.client.Delete("test-data-index", fmt.Sprintf("%d", id))
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error in delete response: %s", res.String())
	}

	return nil
}

func (t *testDataDocumentary) DeletingIndex(ctx context.Context) error {
	res, err := t.client.Indices.Delete([]string{"test-data-index"})
	if err != nil {
		return fmt.Errorf("error deleting index: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error in delete index response: %s", res.String())
	}

	return nil
}

var _ documentaries_interfaces.TestDataDocumentary = &testDataDocumentary{}

func NewTestDataDocumentary(client *elasticsearch.Client) *testDataDocumentary {
	return &testDataDocumentary{
		client: client,
	}
}
