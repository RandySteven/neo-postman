package resolvers

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	resolvers_interfaces "github.com/RandySteven/neo-postman/interfaces/resolvers"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
)

type testDataResolver struct {
	usecase usecases_interfaces.TestDataUsecase
}

func (t *testDataResolver) GetAllTestDatas(ctx context.Context) (result []*responses.TestRecordList, customErr *apperror.CustomError) {
	return t.usecase.GetAllRecords(ctx, nil)
}

var _ resolvers_interfaces.TestDataResolver = &testDataResolver{}

func NewTestDataResolver(usecase usecases_interfaces.TestDataUsecase) *testDataResolver {
	return &testDataResolver{
		usecase: usecase,
	}
}
