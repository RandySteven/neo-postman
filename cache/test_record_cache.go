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

type testRecordCache struct {
	client *redis.Client
}

func (t *testRecordCache) GetMultiData(ctx context.Context) (values []*models.TestRecord, err error) {
	key := "all.test_records"
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

func (t *testRecordCache) SetMultiData(ctx context.Context, values []*models.TestRecord) (err error) {
	key := "all.test_records"
	err = t.client.Set(ctx, key, values, enums.CacheDuration).Err()
	if err != nil {
		return fmt.Errorf("set err: %v", err)
	}
	return nil
}

func (t *testRecordCache) Refresh(ctx context.Context, key string, update any) (value any, err error) {
	if key == "all.test_records" {

	}
	return
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
	if key != "all.test_records" {
		key = fmt.Sprintf("models.test_record:%s", key)
	}
	return utils.Del[models.TestRecord](ctx, t.client, key)
}

func NewTestRecordCache(client *redis.Client) *testRecordCache {
	return &testRecordCache{
		client: client,
	}
}

var _ caches_interfaces.TestRecordCache = &testRecordCache{}
