package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/neo-postman/entities/models"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	"github.com/RandySteven/neo-postman/queries"
	"github.com/RandySteven/neo-postman/utils"
)

type testRecordRepository struct {
	db *sql.DB
}

func (t *testRecordRepository) Save(ctx context.Context, request *models.TestRecord) (result *models.TestRecord, err error) {
	id, err := utils.Save[models.TestRecord](ctx, t.db, queries.InsertTestRecord, &request.TestDataID)
	if err != nil {
		return nil, err
	}
	result = request
	result.ID = *id
	return result, err
}

func (t *testRecordRepository) FindAll(ctx context.Context) (result []*models.TestRecord, err error) {
	return utils.FindAll[models.TestRecord](ctx, t.db, queries.SelectAllTestRecords)
}

func (t *testRecordRepository) FindByID(ctx context.Context, id uint64) (result *models.TestRecord, err error) {
	result = &models.TestRecord{}
	err = utils.FindByID[models.TestRecord](ctx, t.db, queries.SelectTestRecordById, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *testRecordRepository) Update(ctx context.Context, request *models.TestRecord) (result *models.TestRecord, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *testRecordRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.TestRecordRepository = &testRecordRepository{}

func NewTestRecordRepository(db *sql.DB) *testRecordRepository {
	return &testRecordRepository{
		db: db,
	}
}
