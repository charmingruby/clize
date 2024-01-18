package database

import (
	"context"
	"fmt"

	"github.com/charmingruby/clize/config"
	rdb "github.com/go-redis/redis/v8"
)

func Connect(cfg *config.Config) (*rdb.Client, error) {
	ctx := context.Background()

	connectionString := fmt.Sprintf("rediss://default:%s@us1-ruling-gelding-37317.upstash.io:37317", cfg.Redis.Password)

	opt, err := rdb.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}

	redisClient := rdb.NewClient(opt)
	connErr := redisClient.Ping(ctx)
	if connErr.Err() != nil {
		return nil, connErr.Err()
	}

	return redisClient, nil
}
