package limiter

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/senadi/limiter/caches"
)

// Rate - Contains information about remaiming number of requests
// and the reset time.
type Rate struct {
	Remaining int       `json:"remaining"`
	ResetTime time.Time `json:"reset_time"`
}

// RateLimiter - Container about limit/time rate limit requirements
// and chache implementation that is used to achive the same.
type RateLimiter struct {
	Limit    int
	Duration time.Duration
	Cache    caches.RateCache
}

// NewRateLimiter - Creates new RateLimiter object
func NewRateLimiter(limit int, duration time.Duration, cache caches.RateCache) *RateLimiter {
	return &RateLimiter{limit, duration, cache}
}

// Set - RateCache Set implementation - Puts the ip:Rate pair in cache.
func (rl *RateLimiter) Set(ip string, rate Rate) error {
	entry, err := json.Marshal(rate)
	if err != nil {
		log.Error("Error occured during rate limit set")
		return err
	}
	rl.Cache.Set(ip, string(entry[:]), rl.Duration)
	return nil
}

// Get - RateCache Get implementation - Fetch the ip:Rate pair from
func (rl *RateLimiter) Get(ip string) (*Rate, error) {
	var rate *Rate
	entry := rl.Cache.Get(ip)
	if entry == "" {
		return nil, nil
	}
	err := json.Unmarshal([]byte(entry), &rate)
	if err != nil {
		log.Error("Error occured during rate limit fetch")
		return nil, err
	}
	return rate, nil
}
