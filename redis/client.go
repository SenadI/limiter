package redis

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// NewRedisClient - Creates a new redis client
func NewRedisClient(address string, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.WithFields(log.Fields{"address": address, "password": password, "db": db}).Error("Failed to connect to redis!")
		return nil
	}
	return client
}
