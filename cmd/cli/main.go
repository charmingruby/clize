package main

import (
	"log"

	"github.com/charmingruby/clize/config"
	"github.com/charmingruby/clize/internal/database/redis"
)

func main() {
	// Load environment variables
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	redis.Connect(cfg)
}
