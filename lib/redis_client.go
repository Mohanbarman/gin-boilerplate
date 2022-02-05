package lib

import (
	"context"
	"errors"
	"fmt"
	"time"

	"example.com/config"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Ctx    context.Context
	Config *config.RedisConfig
	DB     *redis.Client
	Prefix string
}

func (client *RedisClient) buildKey(key string) string {
	return fmt.Sprintf("%s::::%s", client.Prefix, key)
}

func (client *RedisClient) Get(key string) (value interface{}, err error) {
	value, err = client.DB.Get(client.Ctx, client.buildKey(key)).Result()
	if err != nil {
		err = errors.New("key doesn't exists")
	}
	return
}

func (client *RedisClient) Set(key string, value interface{}, exp time.Duration) (err error) {
	err = client.DB.Set(client.Ctx, client.buildKey(key), value, exp).Err()
	return
}

func (client *RedisClient) Remove(key string) (err error) {
	err = client.DB.Del(client.Ctx, client.buildKey(key)).Err()
	return
}

func GetRedisClient(prefix string, config *config.RedisConfig) *RedisClient {
	return &RedisClient{
		Ctx:    context.Background(),
		Config: config,
		Prefix: prefix,
		DB: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
			Password: config.Password,
			DB:       config.Database,
		}),
	}
}
