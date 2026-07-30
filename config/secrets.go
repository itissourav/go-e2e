package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Secrets struct {
	ConnStr                 string `json:"ConnStr"`
	JWTSecret               string `json:"JWT_SECRET"`
	JWTExpiryHours          string `json:"JWT_EXPIRY_HOURS"`
	RefreshTokenExpiryHours string `json:"REFRESH_TOKEN_EXPIRY_HOURS"`
}

func LoadSecrets() error {
	// Load AWS configuration
	cfg, err := awsConfig.LoadDefaultConfig(context.Background())
	if err != nil {
		return fmt.Errorf("unable to load AWS config: %w", err)
	}

	// Create Secrets Manager client
	client := secretsmanager.NewFromConfig(cfg)

	// Read secret
	result, err := client.GetSecretValue(
		context.Background(),
		&secretsmanager.GetSecretValueInput{
			SecretId: aws.String("go-backend-prod"),
		},
	)
	if err != nil {
		return fmt.Errorf("unable to read secret: %w", err)
	}

	log.Println("Raw Secret:", aws.ToString(result.SecretString))

	var secret Secrets

	err = json.Unmarshal([]byte(*result.SecretString), &secret)
	if err != nil {
		return fmt.Errorf("unable to parse secret: %w", err)
	}
	log.Println(secret.ConnStr, "secret")

	// Populate environment variables
	os.Setenv("ConnStr", secret.ConnStr)
	os.Setenv("JWT_SECRET", secret.JWTSecret)
	os.Setenv("JWT_EXPIRY_HOURS", secret.JWTExpiryHours)
	os.Setenv("REFRESH_TOKEN_EXPIRY_HOURS", secret.RefreshTokenExpiryHours)
	log.Println("Env ConnStr:", os.Getenv("ConnStr"))

	return nil
}
