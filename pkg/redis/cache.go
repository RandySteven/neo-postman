package redis

import (
	"context"
	"fmt"
	"github.com/RandySteven/neo-postman/cache"
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	"github.com/RandySteven/neo-postman/pkg/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type RedisClient struct {
	TestDataCache caches_interfaces.TestDataCache
	client        *redis.Client
}

func NewRedis(config *config.Config) (*RedisClient, error) {
	addr := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	log.Println("connecting to redis : ", addr)
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		MinIdleConns: config.Redis.MinIddleConns,
		PoolSize:     config.Redis.PoolSize,
		PoolTimeout:  time.Duration(config.Redis.PoolTimeout) * time.Second,
		DB:           0,
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

func (r *RedisClient) ClearCache(ctx context.Context) error {
	err := r.client.FlushAll(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to clear cache: %w", err)
	}
	return nil
}
