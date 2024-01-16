package main

import (
	"log"

	"github.com/charmingruby/clize/config"
	rdb "github.com/charmingruby/clize/internal/database/redis"
	"github.com/charmingruby/clize/internal/domain/apps"
	repository "github.com/charmingruby/clize/internal/repository/redis"
)

func main() {
	// Load environment variables
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Redis Connection
	redisClient, err := rdb.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the repositories
	appRepo := repository.NewRedisAppRepository(redisClient)

	// Initialize the services
	apps.NewAppService(appRepo)
}
