package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
)

type ApiCollectionUsecase interface {
	UploadAPICollection(ctx context.Context, request *requests.APICollectionRequest) (result *responses.APICollectionCreateResponse, customErr *apperror.CustomError)
	GetAllCollections(ctx context.Context) (results []*responses.APICollectionListResponse, customErr *apperror.CustomError)
	GetCollectionDetail(ctx context.Context, id uint64) (result *responses.APICollectionDetailResponse, customErr *apperror.CustomError)
	DeleteCollection(ctx context.Context, id uint64) (customErr *apperror.CustomError)
}
