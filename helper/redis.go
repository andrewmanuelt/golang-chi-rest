package helper

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type redisConfig interface {
	GetRedis(string) string
	SetRedis(string, string) error
}

type redisConfigImpl struct {
	context context.Context
	rclient redis.Client
}

// GetRedis implements redisConfig
func (c redisConfigImpl) GetRedis(key string) string {
	redis := c.rclient

	value, err := redis.Get(c.context, key).Result()

	if err != nil {
		return ""
	}

	return value
}

// SetRedis implements redisConfig
func (c redisConfigImpl) SetRedis(key string, value string) error {
	redis := c.rclient

	err := redis.Set(c.context, key, value, 0).Err()

	return err
}

func NewRedis(context context.Context, rclient redis.Client) redisConfig {
	return redisConfigImpl{
		context: context,
		rclient: rclient,
	}
}
