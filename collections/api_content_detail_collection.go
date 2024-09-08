package collections

import (
	"context"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/enums"
	collections_interfaces "github.com/RandySteven/neo-postman/interfaces/collections"
	"github.com/RandySteven/neo-postman/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type apiContentDetailCollection struct {
	collection *mongo.Collection
}

func (a *apiContentDetailCollection) Store(ctx context.Context, entity *models.ApiContentDetail) (result *models.ApiContentDetail, err error) {
	return utils.Store[models.ApiContentDetail](ctx, a.collection, entity)
}

func (a *apiContentDetailCollection) FindAll(ctx context.Context) (result []*models.ApiContentDetail, err error) {
	return utils.Find[models.ApiContentDetail](ctx, a.collection)
}

func (a *apiContentDetailCollection) FindById(ctx context.Context, id primitive.ObjectID) (result *models.ApiContentDetail, err error) {
	return utils.FindID[models.ApiContentDetail](ctx, a.collection, id)
}

func (a *apiContentDetailCollection) DeleteById(ctx context.Context, id primitive.ObjectID) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *apiContentDetailCollection) UpdateById(ctx context.Context, entity *models.ApiContentDetail) (result *models.ApiContentDetail, err error) {
	//TODO implement me
	panic("implement me")
}

var _ collections_interfaces.ApiContentDetailCollection = &apiContentDetailCollection{}

func NewApiContentDetailCollection(mongoDb *mongo.Database) *apiContentDetailCollection {
	return &apiContentDetailCollection{
		collection: mongoDb.Collection(enums.ApiContentDetail),
	}
}
