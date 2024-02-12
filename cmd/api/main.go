package main

import (
	"fmt"
	"log"

	"github.com/charmingruby/clize/config"
	"github.com/charmingruby/clize/internal"

	rdb "github.com/charmingruby/clize/pkg/database/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Redis Connection
	rc, err := rdb.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Services
	appService, err := internal.NewService(rc)
	if err != nil {
		log.Fatal(err)
	}

	// Server
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Handlers
	r, err = internal.NewHTTPService(r, appService)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Run(":" + cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server is running on %s", cfg.Server.Port)
}
