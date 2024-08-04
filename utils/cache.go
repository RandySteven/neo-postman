package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RandySteven/neo-postman/enums"
	"github.com/go-redis/redis/v8"
)

func set(ctx context.Context, client *redis.Client, key string, value interface{}) (err error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}
	return client.Set(ctx, key, jsonData, enums.CacheDuration).Err()
}

func get(ctx context.Context, client *redis.Client, key string) (value interface{}, err error) {
	val, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("get err: %v", err)
	}
	err = json.Unmarshal(val, &value)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err: %v", err)
	}
	return value, nil
}

func Set[T any](ctx context.Context, redis *redis.Client, key string, value *T) (err error) {
	return set(ctx, redis, key, value)
}

func Get[T any](ctx context.Context, redis *redis.Client, key string) (value *T, err error) {
	val, err := get(ctx, redis, key)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	return val.(*T), nil
}

func SetMultiple[T any](ctx context.Context, redis *redis.Client, key string, value []*T) (err error) {
	return set(ctx, redis, key, value)
}

func GetMultiple[T any](ctx context.Context, redis *redis.Client, key string) (value []*T, err error) {
	val, err := get(ctx, redis, key)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	return val.([]*T), nil
}

func Del[T any](ctx context.Context, redis *redis.Client, key string) (err error) {
	return redis.Del(ctx, key).Err()
}
