package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-api-test/entities/models"
	repositories_interfaces "go-api-test/interfaces/repositories"
	"go-api-test/queries"
	"go-api-test/utils"
)

type testDataRepository struct {
	db *sql.DB
}

func jsonString(request map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (t *testDataRepository) Save(ctx context.Context, request *models.TestData) (result *models.TestData, err error) {
	requestHeaderStr, err := jsonString(request.RequestHeader)
	if err != nil {
		return nil, err
	}
	requestBodyStr, err := jsonString(request.RequestBody)
	if err != nil {
		return nil, err
	}
	expectedResponse, err := jsonString(request.ExpectedResponse)
	if err != nil {
		return nil, err
	}
	actualResponse, err := jsonString(request.ActualResponse)
	if err != nil {
		return nil, err
	}
	id, err := utils.Save[models.TestData](ctx, t.db, queries.InsertTestData,
		&request.Method, &request.URI, &request.Description, &requestHeaderStr,
		&requestBodyStr,
		&request.ExpectedResponseCode, &expectedResponse,
		&request.ActualResponseCode, &actualResponse, &request.ResultStatus)
	if err != nil {
		return nil, err
	}
	result = request
	result.ID = *id
	return result, nil
}

func (t *testDataRepository) FindAll(ctx context.Context) (result []*models.TestData, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *testDataRepository) FindByID(ctx context.Context, id *uint64) (result *models.TestData, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *testDataRepository) Update(ctx context.Context, request *models.TestData) (result *models.TestData, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *testDataRepository) Delete(ctx context.Context, id *uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.TestDataRepository = &testDataRepository{}

func NewTestDataRepository(db *sql.DB) *testDataRepository {
	return &testDataRepository{
		db: db,
	}
}
