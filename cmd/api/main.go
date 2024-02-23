package main

import (
	"fmt"
	"log"

	"github.com/charmingruby/clize/config"
	"github.com/charmingruby/clize/internal"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	rdb "github.com/charmingruby/clize/pkg/database/redis"
	"github.com/gin-contrib/cors"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found.")
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
	log.Println("Creating services...")
	appService, err := internal.NewService(rc)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Services created.")

	// Server
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(
		cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"Origin", "Accept", "Content-Type", "Authorization"},
		}),
	)

	// Handlers
	log.Println("Creating HTTP service...")
	r, err = internal.NewHTTPService(r, appService)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Service created.")

	fmt.Printf("Server is running on %s", cfg.Server.Port)
	err = r.Run(":" + cfg.Server.Port)
	if err != nil {
		log.Fatal(err.Error())
	}

}
