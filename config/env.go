package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMONGOURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error to load your .env file")
	}

	return os.Getenv("MONGOURI")
}
