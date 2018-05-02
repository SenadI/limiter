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
	log "github.com/sirupsen/logrus"
)

func init() {
	rootCmd.AddCommand(rlMemCmd)
}

var rlMemCmd = &cobra.Command{
	Use:   "rlmem",
	Short: "Run the HTTP server with in-memory rate limiting",
	Long:  fmt.Sprintf("Run the HTTP server on the provided address with in-memory rate limiting"),
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: Extract configuration structure for this
		address := viper.GetString("address")
		limit := viper.GetInt("limit")
		duration := viper.GetDuration("duration")

		inMemoryCache := caches.NewInMemoryCache(duration)
		inMemoryRateLimiter := limiter.NewRateLimiter(limit, duration, &inMemoryCache)

		r := mux.NewRouter()
		r.Handle("/greet/me", middleware.RateLimiterMiddleware(inMemoryRateLimiter, limiterHttp.NewGreetingsHandler()))
		log.WithField("address", address).Info("Server started...")
		log.Info(http.ListenAndServe(address, r))
	},
}
