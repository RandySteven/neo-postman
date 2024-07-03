package usecases_interfaces

import (
	"context"
	"go-api-test/apperror"
	"go-api-test/entities/payloads/requests"
	"go-api-test/entities/payloads/responses"
)

type TestDataUsecase interface {
	CreateAPITest(ctx context.Context, request *requests.TestDataRequest) (result *responses.TestDataResponse, customErr *apperror.CustomError)
}
