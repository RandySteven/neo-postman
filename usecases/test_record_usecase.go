package usecases

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
)

type testRecordUsecase struct {
	testDataRepo   repositories_interfaces.TestDataRepository
	testRecordRepo repositories_interfaces.TestRecordRepository
}

func (t *testRecordUsecase) AutoSaveTestRecord(ctx context.Context) (err error) {
	return
}

func (t *testRecordUsecase) CreateTestRecord(ctx context.Context, request *requests.TestRecordRequest) (customErr *apperror.CustomError) {
	return
}

func (t *testRecordUsecase) GetAllTestRecords(ctx context.Context) (result []*responses.TestRecordListResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (t *testRecordUsecase) GetTestRecordDetail(ctx context.Context, id uint64) (result *responses.TestRecordDetailResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

var _ usecases_interfaces.TestRecordUseCase = &testRecordUsecase{}

func NewTestRecordUsecase(
	testDataRepo repositories_interfaces.TestDataRepository,
	testRecordRepo repositories_interfaces.TestRecordRepository,
) *testRecordUsecase {
	return &testRecordUsecase{
		testDataRepo:   testDataRepo,
		testRecordRepo: testRecordRepo,
	}
}
