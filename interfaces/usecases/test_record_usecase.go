package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
)

type TestRecordUseCase interface {
	AutoSaveTestRecord(ctx context.Context) (err error)
	CreateTestRecord(ctx context.Context, request *requests.TestRecordRequest) (result *responses.TestRecordCreateResponse, customErr *apperror.CustomError)
	GetAllTestRecords(ctx context.Context) (result []*responses.TestRecordListResponse, customErr *apperror.CustomError)
	GetTestRecordDetail(ctx context.Context, id uint64) (result *responses.TestRecordDetailResponse, customErr *apperror.CustomError)
}
