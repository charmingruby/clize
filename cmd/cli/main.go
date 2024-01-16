package main

import (
	"log"

	"github.com/charmingruby/clize/config"
)

func main() {
	// Load environment variables
	_, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
}
