package main

import (
	"fmt"
	"log"

	"github.com/charmingruby/clize/config"
	"github.com/charmingruby/clize/internal/auth"
	rdb "github.com/charmingruby/clize/internal/database/redis"
	"github.com/charmingruby/clize/internal/domain/apps"
	repository "github.com/charmingruby/clize/internal/repository/redis"
	"github.com/charmingruby/clize/internal/transport/rest"
)

const (
	ApiPort = "8080"
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

	// Authenticator
	a, err := auth.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Server
	r := rest.New(a)
	err = r.Run(":" + ApiPort)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server is running on %s", ApiPort)
}
