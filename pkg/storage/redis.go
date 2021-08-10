package storage

import (
	"strconv"

	"github.com/chiehting/apiGo-template/pkg/config"
	"github.com/go-redis/redis/v8"
)

// Redis is connection to redis
var Redis = redisConnection()

func redisConnection() *redis.Client {
	cfg := config.GetCache()
	DB, _ := strconv.Atoi(cfg.DB)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       DB,
	})

	return redisClient
}
