package caches

import (
	"time"
)

// RateCache - abstracts the cache backend. Any cache struct that implements
// this interface can be used for rate limiting.
type RateCache interface {
	Set(key string, object string, duration time.Duration)
	Get(key string) string
}
