package resolvers_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
)

type TestDataResolver interface {
	GetAllTestDatas(ctx context.Context) (result []*responses.TestRecordList, customErr *apperror.CustomError)
}
