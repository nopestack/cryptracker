package config

import (
	"log"
	"os"
	"fmt"

	"github.com/alexsasharegan/dotenv"
)

func init() {
	verbose := os.Getenv("DEBUG")
	if verbose == "true" {
		fmt.Println("Reading from .env")
	}

	err := dotenv.Load()

	if err != nil {
		log.Fatal("Error loading configuration files")
	}
}
