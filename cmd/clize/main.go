package main

import (
	"fmt"
	"log"

	"github.com/charmingruby/clize/config"
	"github.com/charmingruby/clize/internal/application"
	"github.com/charmingruby/clize/internal/auth"
	"github.com/charmingruby/clize/internal/auth/domain"
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
	applicationSvc, err := application.NewService(rc)
	if err != nil {
		log.Fatal(err)
	}

	// Server
	r := gin.Default()

	// Handlers
	_, err = application.NewHTTPService(r, applicationSvc)
	if err != nil {
		log.Fatal(err)
	}

	// Auth Handler
	// Authenticator
	authenticator, err := domain.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	_, err = auth.NewHTTPService(r, authenticator)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Run(":" + ApiPort)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server is running on %s", ApiPort)
}