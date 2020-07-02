package client

import (
	"github.com/go-redis/redis"
)

type RedisClient struct {
}

var (
	redisClient *redis.Client
)

func GetRedisClient() (c *redis.Client) {
	return redisClient
}

func init() {
	var (
		err error
	)

	redisClient = redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:           "",
		DB:                 0,
	})

	if _, err = redisClient.Ping().Result(); err != nil {
		return
	}
	return
}