package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var cache *redis.Client
var ctx = context.Background()

const (
	CACHE_HOST = "localhost:6379"
	CACHE_PW   = ""
)

func InitCache() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     CACHE_HOST,
		Password: CACHE_PW,
		DB:       0,
	})
	cache = rdb
}

func Get(key string) string {
	val, err := cache.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("cache miss")
		return ""
	}
	return val
}

func Set(key string, val string) string {
	err := cache.Set(ctx, key, val, 0).Err()
	if err != nil {
		panic(err)
	}
	return val
}
