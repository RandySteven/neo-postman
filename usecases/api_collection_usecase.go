package usecases

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
)

type apiCollectionUsecase struct {
	apiRepository      repositories_interfaces.ApiRepository
	apiCollectionCache caches_interfaces.ApiCollectionCache
}

func (a *apiCollectionUsecase) UploadAPICollection(ctx context.Context, request *requests.APICollectionRequest) (result *responses.APICollectionCreateResponse, customErr *apperror.CustomError) {
	panic("implement me")
}

func (a *apiCollectionUsecase) GetAllCollections(ctx context.Context) (results []*responses.APICollectionListResponse, customErr *apperror.CustomError) {
	apis, err := a.apiCollectionCache.GetMultiData(ctx)
	if err != nil {
		apis, err = a.apiRepository.FindAll(ctx)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get api`, err)
		}
	}

	for _, api := range apis {
		results = append(results, &responses.APICollectionListResponse{
			ID:        api.ID,
			Title:     api.Title,
			CreatedAt: api.CreatedAt,
		})
	}
	return results, nil
}

func (a *apiCollectionUsecase) GetCollectionDetail(ctx context.Context, id uint64) (result *responses.APICollectionDetailResponse, customErr *apperror.CustomError) {
	panic("implement me")
}

func (a *apiCollectionUsecase) DeleteCollection(ctx context.Context, id uint64) (customErr *apperror.CustomError) {
	panic("implement me")
}

var _ usecases_interfaces.ApiCollectionUsecase = &apiCollectionUsecase{}

func NewApiCollectionUsecase(apiRepository repositories_interfaces.ApiRepository) *apiCollectionUsecase {
	return &apiCollectionUsecase{apiRepository: apiRepository}
}
