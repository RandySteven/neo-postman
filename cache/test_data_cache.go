package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/enums"
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	"github.com/go-redis/redis/v8"
)

type testDataCache struct {
	client *redis.Client
}

func (t *testDataCache) Set(ctx context.Context, key string, value *models.TestData) (err error) {
	key = fmt.Sprintf("models.test_data.%s", key)
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal test data: %w", err)
	}
	return t.client.Set(ctx, key, jsonData, enums.CacheDuration).Err()
}

func (t *testDataCache) Get(ctx context.Context, key string) (value *models.TestData, err error) {
	key = fmt.Sprintf("models.test_data.%s", key)
	val, err := t.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to get record: %w", err)
	}
	var testData models.TestData
	if err = json.Unmarshal(val, &testData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal test data: %w", err)
	}
	return &testData, nil
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
