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

func (t *testDataDocumentary) MakeAnIndex(ctx context.Context, document *models.TestData) (err error) {
	_, err = t.client.Indices.Create("test-data-index")
	if err != nil {
		return err
	}
	data, _ := json.Marshal(document)
	_, err = t.client.Index("my_index", bytes.NewReader(data))
	if err != nil {
		return err
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

func (t *testDataDocumentary) SearchDocument(ctx context.Context, query string) (result *models.TestData, err error) {
	//TODO implement me
	bodyQry := fmt.Sprintf(`{
		"query": {
			"multi_match": {
				"query": "%s",
				"fields": [
					"id",
				 	"host", "method", "uri", "description", 
				 	"request_header", "request_body", "actual_response_code", "expected_response_code",
					"result_status"
				]
			}
		}
	}`, query)
	log.Println(bodyQry)

	res, err := t.client.Search(
		t.client.Search.WithContext(ctx),
		t.client.Search.WithIndex("test-data-index"),
		t.client.Search.WithBody(bytes.NewBufferString(bodyQry)),
		t.client.Search.WithPretty())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf(res.String())
	}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return
}

func (t *testDataDocumentary) DeletingDocument(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t *testDataDocumentary) DeletingIndex(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}

var _ documentaries_interfaces.TestDataDocumentary = &testDataDocumentary{}

func NewTestDataDocumentary(client *elasticsearch.Client) *testDataDocumentary {
	return &testDataDocumentary{
		client: client,
	}
}
