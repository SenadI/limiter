package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/senadi/limiter"
	"github.com/senadi/limiter/caches"
	limiterHttp "github.com/senadi/limiter/http"
	"github.com/senadi/limiter/http/middleware"
	"github.com/senadi/limiter/redis"
	log "github.com/sirupsen/logrus"
)

func init() {
	rootCmd.AddCommand(rlRedisCmd)
}

var rlRedisCmd = &cobra.Command{
	Use:   "rlredis",
	Short: "Run the HTTP server with redis backend for rate limiting",
	Long:  fmt.Sprintf("Run the HTTP server on the provided address wtih redis backed for rate limiting"),
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: Extract configuration structure for this
		address := viper.GetString("address")
		limit := viper.GetInt("limit")
		duration := viper.GetDuration("duration")
		redisAddress := viper.GetString("redis.address")
		redisPassword := viper.GetString("redis.password")
		redisDatabase := viper.GetInt("redis.db")

		redisClient := redis.NewRedisClient(redisAddress, redisPassword, redisDatabase)
		redisCache := caches.NewRedisCache(duration, redisClient)
		redisRateLimiter := limiter.NewRateLimiter(limit, duration, &redisCache)

		r := mux.NewRouter()
		r.Handle("/greet/me/redis", middleware.RateLimiterMiddleware(redisRateLimiter, limiterHttp.NewGreetingsHandler()))

		log.WithField("address", address).Info("Server started...")
		log.Info(http.ListenAndServe(address, r))
	},
}
