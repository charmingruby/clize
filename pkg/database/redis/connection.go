package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/charmingruby/clize/config"
	rdb "github.com/go-redis/redis/v8"
)

func Connect(cfg *config.Config) (*rdb.Client, error) {
	log.Println("Connection to Redis Upstash...")

	ctx := context.Background()

	connectionString := fmt.Sprintf("rediss://default:%s@us1-ruling-gelding-37317.upstash.io:37317", cfg.Redis.Password)

	opt, err := rdb.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}

	redisClient := rdb.NewClient(opt)
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Redis Upstash.")

	return redisClient, nil
}
