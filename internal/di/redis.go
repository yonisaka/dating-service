package di

import (
	"fmt"
	"log"

	"github.com/yonisaka/dating-service/pkg/di"
	"github.com/yonisaka/dating-service/pkg/kvs"
	"github.com/yonisaka/dating-service/pkg/kvs/redis"
)

// GetRedis get the Redis KVS client.
func GetRedis() kvs.Client {
	cfg := GetConfig()
	r, err := redis.New(
		redis.WithAddr(fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)),
		redis.WithDB(cfg.Redis.DB),
	)

	if err != nil {
		log.Fatal(err)
	}

	di.RegisterCloser("RedisConnection", r)

	return r
}
