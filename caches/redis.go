package caches

import (
	"time"

	"github.com/go-redis/redis"
)

// RedisCache - Simple cache structure that is uses redis rate limiting.
type RedisCache struct {
	DefaultExpiry time.Duration
	Client        *redis.Client
}

// NewRedisCache - Creates a new redis cache to be consumed by rate-limiter.
func NewRedisCache(defaultExpiry time.Duration, client *redis.Client) RedisCache {
	return RedisCache{DefaultExpiry: defaultExpiry, Client: client}
}

// Set - Writes object for key to redis.
func (c *RedisCache) Set(key string, object string, duration time.Duration) {
	c.Client.Set(key, object, duration)
}

// Get - Retrieves object for key pair from redis.
func (c *RedisCache) Get(key string) string {
	entry := c.Client.Get(key)
	if entry != nil {
		return entry.Val()
	}
	return ""
}
