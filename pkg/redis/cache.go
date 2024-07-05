package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedis() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisClient{
		client: client,
	}
}

func (client *RedisClient) Client() *redis.Client {
	return client.client
}

func (client *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return client.client.Set(ctx, key, value, expiration).Err()
}

func (client *RedisClient) Get(ctx context.Context, key string) (interface{}, error) {
	return client.client.Get(ctx, key).Result()
}

func (r *RedisClient) Ping(ctx context.Context) error {
	result, err := r.client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
