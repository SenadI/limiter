package caches

import (
	"time"

	goCache "github.com/patrickmn/go-cache"
)

// InMemoryCache - Simple cache structure that is used for in-memory rate limiting.
type InMemoryCache struct {
	DefaultExpiry time.Duration
	Cache         *goCache.Cache
}

// NewInMemoryCache - Creates a new in-memory cache to be consumed by rate-limiter.
func NewInMemoryCache(defaultExpiry time.Duration) InMemoryCache {
	c := goCache.New(defaultExpiry, time.Millisecond*100)
	return InMemoryCache{DefaultExpiry: defaultExpiry, Cache: c}

}

// Set - Writes object for key to go-cache.
func (c *InMemoryCache) Set(key string, object string, duration time.Duration) {
	c.Cache.Set(key, object, duration)
}

// Get - Retrieves object for key pair from go-cache.
func (c *InMemoryCache) Get(key string) string {
	entry, ok := c.Cache.Get(key)
	if ok {
		return entry.(string)
	}
	return ""
}
