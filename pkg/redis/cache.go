package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedis() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisClient{
		client: client,
	}
}

func (client *RedisClient) Client() *redis.Client {
	return client.client
}

func (r *RedisClient) Ping(ctx context.Context) error {
	_, err := r.client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
