package cache

import (
	"context"
	"fmt"
	"github.com/RandySteven/neo-postman/entities/models"
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	"github.com/RandySteven/neo-postman/utils"
	"github.com/go-redis/redis/v8"
)

type apiCollectionCache struct {
	client *redis.Client
}

func (a apiCollectionCache) Set(ctx context.Context, key string, value *models.Api) (err error) {
	key = fmt.Sprintf("models.api.%s", key)
	return utils.Set[models.Api](ctx, a.client, key, value)
}

func (a apiCollectionCache) Get(ctx context.Context, key string) (value *models.Api, err error) {
	key = fmt.Sprintf("models.api.%s", key)
	return utils.Get[models.Api](ctx, a.client, key)
}

func (a apiCollectionCache) SetMultiData(ctx context.Context, values []*models.Api) (err error) {
	key := "all.apis"
	return utils.SetMultiple[models.Api](ctx, a.client, key, values)
}

func (a apiCollectionCache) GetMultiData(ctx context.Context) (values []*models.Api, err error) {
	key := "all.apis"
	return utils.GetMultiple[models.Api](ctx, a.client, key)
}

func (a apiCollectionCache) Refresh(ctx context.Context, key string, update any) (value any, err error) {
	//TODO implement me
	panic("implement me")
}

func (a apiCollectionCache) Del(ctx context.Context, key string) (err error) {
	return utils.Del[models.Api](ctx, a.client, key)
}

var _ caches_interfaces.ApiCollectionCache = &apiCollectionCache{}

func NewApiCollectionCache(client *redis.Client) *apiCollectionCache {
	return &apiCollectionCache{
		client: client,
	}
}
