package redis

import (
	"context"
	"fmt"

	"rotom/config"

	"github.com/go-redis/redis/v8"
)

var (
	ctx 	= context.Background()
	Redis		*redis.Client
)

func init() {
	r := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB: 0,
	})

	Redis = r

	defer r.Close()
}

func Set(key string, value interface{}) error {
	return Redis.Set(ctx, key, value, 0).Err()
}

func Get(key string) (string, error) {
	return Redis.Get(ctx, key).Result()
}

func Del(key string) error {
	return Redis.Del(ctx, key).Err()
}
