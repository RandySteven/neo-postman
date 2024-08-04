package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RandySteven/neo-postman/enums"
	"github.com/go-redis/redis/v8"
)

func set(ctx context.Context, client *redis.Client, key string, value interface{}) (err error) {
	return
}

func get(ctx context.Context, client *redis.Client, key string) (value interface{}, err error) {
	return
}

func Set[T any](ctx context.Context, redis *redis.Client, key string, value *T) (err error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}
	return redis.Set(ctx, key, jsonData, enums.CacheDuration).Err()
}

func Get[T any](ctx context.Context, redis *redis.Client, key string) (value *T, err error) {
	val, err := redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("get err: %v", err)
	}
	err = json.Unmarshal(val, &value)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err: %v", err)
	}
	return value, nil
}

func SetMultiple[T any](ctx context.Context, redis *redis.Client, key string, value []*T) (err error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}
	return redis.Set(ctx, key, jsonData, enums.CacheDuration).Err()
}

func GetMultiple[T any](ctx context.Context, redis *redis.Client, key string) (value []*T, err error) {
	val, err := redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("get err: %v", err)
	}
	err = json.Unmarshal(val, &value)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err: %v", err)
	}
	return value, nil
}

func Del[T any](ctx context.Context, redis *redis.Client, key string) (err error) {
	return redis.Del(ctx, key).Err()
}
