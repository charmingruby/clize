package main

import (
	"fmt"
	"log"

	"github.com/charmingruby/clize/config"
	"github.com/charmingruby/clize/internal/app"
	"github.com/charmingruby/clize/internal/auth"
	"github.com/charmingruby/clize/internal/common"

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
	appService, err := app.NewService(rc)
	if err != nil {
		log.Fatal(err)
	}

	authSvc, err := auth.NewService(rc)
	if err != nil {
		log.Fatal(err)
	}

	// Server
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Handlers
	r, err = app.NewHTTPService(r, appService)
	if err != nil {
		log.Fatal(err)
	}

	r, err = auth.NewHTTPService(r, authSvc)
	if err != nil {
		log.Fatal(err)
	}

	r, err = common.NewHTTPService(r)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Run(":" + cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server is running on %s", cfg.Server.Port)
}
