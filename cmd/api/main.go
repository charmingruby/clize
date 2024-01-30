package main

import (
	"fmt"
	"log"

	"github.com/charmingruby/clize/config"
	"github.com/charmingruby/clize/internal/application"
	"github.com/charmingruby/clize/internal/common"
	rdb "github.com/charmingruby/clize/pkg/database/redis"
	"github.com/gin-gonic/gin"
)

var ApiPort = "8080"

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
	services, err := application.NewService(rc)
	if err != nil {
		log.Fatal(err)
	}

	// Server
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Handlers
	r, err = application.NewHTTPService(r, services)
	if err != nil {
		log.Fatal(err)
	}

	r, err = common.NewHTTPService(r)
	if err != nil {
		log.Fatal(err)
	}

	// Auth Handler

	err = r.Run(":" + ApiPort)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server is running on %s", ApiPort)
}
