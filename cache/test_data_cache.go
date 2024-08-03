package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/enums"
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	"github.com/RandySteven/neo-postman/utils"
	"github.com/go-redis/redis/v8"
)

type testDataCache struct {
	client *redis.Client
}

func (t *testDataCache) SetMultiData(ctx context.Context, values []*models.TestData) error {
	key := "all.test_datas"
	jsonData, err := json.Marshal(values)
	if err != nil {
		return fmt.Errorf("failed to marshal test data: %w", err)
	}
	err = t.client.Set(ctx, key, jsonData, enums.CacheDuration).Err()
	if err != nil {
		return fmt.Errorf("set err: %v", err)
	}
	return nil
}

func (t *testDataCache) GetMultiData(ctx context.Context) (values []*models.TestData, err error) {
	key := "all.test_datas"
	val, err := t.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("get err: %v", err)
	}
	err = json.Unmarshal(val, &values)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err: %v", err)
	}
	return values, nil
}

func (t *testDataCache) Refresh(ctx context.Context, key string, update any) (value any, err error) {
	//TODO implement me
	panic("implement me")
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
