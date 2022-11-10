package cache

import (
	"context"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"github.com/nitishm/go-rejson"
	"time"
)

type RedisCache struct {
	rdb *redis.Client
	rh  *rejson.Handler
}

func (e *RedisCache) Set(ctx context.Context, key string, val string) error {
	return e.rdb.Set(key, val, 30*time.Second).Err()
}

func (e *RedisCache) Get(ctx context.Context, key string) (val string, err error) {
	val, err = e.rdb.Get(key).Result()
	if err == redis.Nil {
		err = NotExistError
	}
	return val, err
}

func New(rcl *redis.Client) *RedisCache {
	return &RedisCache{rdb: rcl}
}
