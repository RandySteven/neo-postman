package redis

import (
	"context"
	"github.com/RandySteven/neo-postman/cache"
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	"github.com/RandySteven/neo-postman/pkg/config"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisClient struct {
	TestDataCache caches_interfaces.TestDataCache
	client        *redis.Client
}

func NewRedis(config *config.Config) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisClient{
		TestDataCache: cache.NewTestDataCache(client),
		client:        client,
	}, nil
}

func (r *RedisClient) Ping(ctx context.Context) error {
	result, err := r.client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
