package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmingruby/clize/config"
	"github.com/charmingruby/clize/internal"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	rdb "github.com/charmingruby/clize/pkg/database/redis"
	"github.com/gin-contrib/cors"
)

func main() {
	// Load environment variables
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

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
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(
		cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"Origin", "Accept", "Content-Type"},
		}),
	)

	// Handlers
	r, err = internal.NewHTTPService(r, appService)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server is running on %s", cfg.Server.Port)
	err = r.Run(":" + cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

}
