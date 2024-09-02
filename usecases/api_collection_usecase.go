package usecases

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
)

type apiCollectionUsecase struct {
	apiRepository repositories_interfaces.ApiRepository
}

func (a *apiCollectionUsecase) UploadAPICollection(ctx context.Context, request *requests.APICollectionRequest) (result *responses.APICollectionCreateResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (a *apiCollectionUsecase) GetAllCollections(ctx context.Context) (results []*responses.APICollectionListResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (a *apiCollectionUsecase) GetCollectionDetail(ctx context.Context, id uint64) (result *responses.APICollectionDetailResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (a *apiCollectionUsecase) DeleteCollection(ctx context.Context, id uint64) (customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

var _ usecases_interfaces.ApiCollectionUsecase = &apiCollectionUsecase{}

func NewApiCollectionUsecase(apiRepository repositories_interfaces.ApiRepository) *apiCollectionUsecase {
	return &apiCollectionUsecase{apiRepository: apiRepository}
}
