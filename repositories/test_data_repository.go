package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/neo-postman/entities/models"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	"github.com/RandySteven/neo-postman/queries"
	"github.com/RandySteven/neo-postman/utils"
)

type testDataRepository struct {
	db *sql.DB
}

func (t *testDataRepository) Save(ctx context.Context, request *models.TestData) (result *models.TestData, err error) {
	requestHeaderStr, err := utils.JsonString(request.RequestHeader)
	if err != nil {
		return nil, err
	}
	requestBodyStr, err := utils.JsonString(request.RequestBody)
	if err != nil {
		return nil, err
	}
	expectedResponse, err := utils.JsonString(request.ExpectedResponse)
	if err != nil {
		return nil, err
	}
	actualResponse, err := utils.JsonString(request.ActualResponse)
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
