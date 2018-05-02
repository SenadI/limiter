package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Tomasen/realip"
	"github.com/senadi/limiter"
	json "github.com/senadi/limiter/pkg/json"
)

// RateLimiterMiddleware - The actual logic behind rate-limiting is implemented here.
func RateLimiterMiddleware(l *limiter.RateLimiter, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// first we get the IP
		var ip = realip.FromRequest(r)
		// get the rate limit entry from the cache
		rate, err := l.Get(ip)
		if err != nil {
			writeError(w)
		}
		var currentTime = time.Now()
		var resetTime = currentTime.Add(l.Duration)
		// if this is the first request set fresh rate limit
		if rate == nil {
			err = l.Set(ip, limiter.Rate{Remaining: l.Limit, ResetTime: resetTime})
			if err != nil {
				writeError(w)
				return
			}
			// if the rate has been reached write 429 headers
		} else if rate.Remaining == 0 && rate.ResetTime.Sub(currentTime) > 0 {
			setHeaders(l.Limit, rate.Remaining, rate.ResetTime, w)
			w.WriteHeader(http.StatusTooManyRequests)
			json.WriteJSON(http.StatusText(http.StatusTooManyRequests), w)
			return
			// if reset time has been reched - set fresh rate limit
		} else if rate.ResetTime.Sub(currentTime) <= 0 {
			err = l.Set(ip, limiter.Rate{Remaining: l.Limit, ResetTime: resetTime})
			if err != nil {
				writeError(w)
				return
			}
			// decrement the remaining rate coutner
		} else {
			rate.Remaining--
			err = l.Set(ip, *rate)
			if err != nil {
				writeError(w)
				return
			}
		}
		// Refresh the entry from cache
		rate, err = l.Get(ip)
		if err != nil {
			writeError(w)
		}
		setHeaders(l.Limit, rate.Remaining, rate.ResetTime, w)
		next.ServeHTTP(w, r)
	}
}

func setHeaders(limit int, remaining int, resetTime time.Time, w http.ResponseWriter) {
	w.Header().Set("X-RateLimit-Limit", strconv.Itoa(limit))
	w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(remaining))
	w.Header().Set("X-RateLimit-Reset", resetTime.String())
}

func writeError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.WriteJSON(http.StatusText(http.StatusInternalServerError), w)
}
