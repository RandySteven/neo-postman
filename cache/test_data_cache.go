package cache

import (
	"context"
	"fmt"
	"github.com/RandySteven/neo-postman/entities/models"
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	"github.com/RandySteven/neo-postman/utils"
	"github.com/go-redis/redis/v8"
)

type testDataCache struct {
	client *redis.Client
}

func (t *testDataCache) Set(ctx context.Context, key string, value *models.TestData) (err error) {
	key = fmt.Sprintf("models.test_data.%s", key)
	return utils.Set[models.TestData](ctx, t.client, key, value)
}

func (t *testDataCache) Get(ctx context.Context, key string) (value *models.TestData, err error) {
	key = fmt.Sprintf("models.test_data.%s", key)
	return utils.Get[models.TestData](ctx, t.client, key)
}

func (t *testDataCache) Del(ctx context.Context, key string) (err error) {
	return utils.Del[models.TestData](ctx, t.client, key)
}

func NewTestDataCache(client *redis.Client) *testDataCache {
	return &testDataCache{
		client: client,
	}
}

var _ caches_interfaces.TestDataCache = &testDataCache{}
