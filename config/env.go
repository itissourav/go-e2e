package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Println(".env file not found")
	}
}
