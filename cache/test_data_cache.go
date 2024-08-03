package cache

import (
	"context"
	"encoding/json"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/enums"
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	"github.com/go-redis/redis/v8"
)

type testDataCache struct {
	client *redis.Client
}

func (t *testDataCache) Set(ctx context.Context, key string, value *models.TestData) (err error) {
	return t.client.Set(ctx, key, value, enums.CacheDuration).Err()
}

func (t *testDataCache) Get(ctx context.Context, key string) (value *models.TestData, ok bool) {
	val, err := t.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, false
	}
	var testData models.TestData
	if err := json.Unmarshal(val, &testData); err != nil {
		return nil, false
	}
	return &testData, true
}

func (t *testDataCache) Del(ctx context.Context, key string) (err error) {
	return t.client.Del(ctx, key).Err()
}

func NewTestDataCache(client *redis.Client) *testDataCache {
	return &testDataCache{
		client: client,
	}
}

var _ caches_interfaces.TestDataCache = &testDataCache{}
