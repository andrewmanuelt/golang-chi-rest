package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func RedisConnect() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr: ConfigField("redis", "host"),
		DB:   10,
	})

	fmt.Println("Connected to redis...")

	return r
}
