package cache

import (
	"context"
	"fmt"
	"skello/internal/config"
	"skello/internal/logger"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func Init() {
	host, port := config.RedisConfig()

	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Get().WithError(err).Fatal("Failed to connect to Redis")
	}
	logger.Get().Info("Connected to Redis successfully")
}

func Get() *redis.Client {
	return redisClient
}

func Close() {
	if redisClient != nil {
		redisClient.Close()
	}
}

// TODO: rest of the prod. setup, check the docs and existing community posts on the matter
// https://redis.io/docs/latest/develop/clients/go/connect/ just in case, start here, work your way down future me!
