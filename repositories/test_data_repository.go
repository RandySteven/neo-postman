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
	id, err := utils.Save[models.TestData](ctx, t.db, queries.InsertTestData,
		&request.Method, &request.URI, &request.Description, &request.RequestHeader,
		&request.RequestBody,
		&request.ExpectedResponseCode, &request.ExpectedResponse,
		&request.ActualResponseCode, &request.ActualResponse, &request.ResultStatus)
	if err != nil {
		return nil, err
	}
	result = request
	result.ID = *id
	return result, nil
}

func (t *testDataRepository) FindAll(ctx context.Context) (result []*models.TestData, err error) {
	return utils.FindAll[models.TestData](ctx, t.db, queries.SelectTestData)
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
