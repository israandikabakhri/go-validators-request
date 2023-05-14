package app

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariabels() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file...")
	}

}
