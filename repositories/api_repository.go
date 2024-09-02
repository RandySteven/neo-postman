package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/neo-postman/entities/models"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	"github.com/RandySteven/neo-postman/queries"
	"github.com/RandySteven/neo-postman/utils"
)

type apiRepository struct {
	db *sql.DB
}

func (a *apiRepository) Save(ctx context.Context, request *models.Api) (result *models.Api, err error) {
	id, err := utils.Save[models.Api](ctx, a.db, queries.InsertApiQuery, &request.Title, &request.Description, &request.ContentFile)
	if err != nil {
		return nil, err
	}
	result = request
	result.ID = *id
	return result, nil
}

func (a *apiRepository) FindAll(ctx context.Context) (result []*models.Api, err error) {
	result, err = utils.FindAll[models.Api](ctx, a.db, queries.SelectApisQuery)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *apiRepository) FindByID(ctx context.Context, id uint64) (result *models.Api, err error) {
	err = utils.FindByID[models.Api](ctx, a.db, queries.SelectApiByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *apiRepository) Update(ctx context.Context, request *models.Api) (result *models.Api, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *apiRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.ApiRepository = &apiRepository{}

func NewAPIRepository(db *sql.DB) *apiRepository {
	return &apiRepository{
		db: db,
	}
}
