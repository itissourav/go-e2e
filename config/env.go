package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	log.Println(env, "env-------")

	switch env {
	case "prod":
		err := LoadSecrets("prod")
		if err != nil {
			log.Fatalf("failed to load secrets: %v", err)
		}
	case "dev":
		err := LoadSecrets("dev")
		if err != nil {
			log.Fatalf("failed to load secrets: %v", err)
		}
	default:
		err := godotenv.Load("local.env")
		if err != nil {
			log.Println("local.env file not found")
		}
	}
}
