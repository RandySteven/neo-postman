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

type testRecordCache struct {
	client *redis.Client
}

func (t *testRecordCache) Set(ctx context.Context, key string, value *models.TestRecord) (err error) {
	key = fmt.Sprintf("models.test_record:%s", key)
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return t.client.Set(ctx, key, jsonData, enums.CacheDuration).Err()
}

func (t *testRecordCache) Get(ctx context.Context, key string) (value *models.TestRecord, err error) {
	key = fmt.Sprintf("models.test_record:%s", key)
	jsonData, err := t.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (t *testRecordCache) Del(ctx context.Context, key string) (err error) {
	key = fmt.Sprintf("models.test_record:%s", key)
	return t.client.Del(ctx, key).Err()
}

func NewTestRecordCache(client *redis.Client) *testRecordCache {
	return &testRecordCache{
		client: client,
	}
}

var _ caches_interfaces.TestRecordCache = &testRecordCache{}
