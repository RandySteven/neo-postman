package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	"github.com/RandySteven/neo-postman/queries"
)

type TestDataUsecase interface {
	CreateAPITest(ctx context.Context, request *requests.TestDataRequest) (result *responses.TestDataResponse, customErr *apperror.CustomError)
	GetAllRecords(ctx context.Context, query *queries.QueryParam) (result []*responses.TestRecordList, customErr *apperror.CustomError)
	GetRecord(ctx context.Context, id uint64) (result *responses.TestDataDetail, customErr *apperror.CustomError)
	SaveRecord(ctx context.Context, id uint64) (result string, customErr *apperror.CustomError)
	UnsavedRecord(ctx context.Context, id uint64) (result string, customErr *apperror.CustomError)

	AutoDeleteUnsavedRecord(ctx context.Context) (err error)
}
