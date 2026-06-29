package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Email     string `json:"email"`
	UserID    int64  `json:"user_id"`
	CreatedAt int64  `json:"created_at"`

	jwt.RegisteredClaims
}

func GenerateJWT(userID int64, email string) (string, error) {

	createdAt := time.Now()
	expiresAt := createdAt.Add(24 * time.Hour)

	claims := JWTClaims{
		Email:     email,
		UserID:    userID,
		CreatedAt: createdAt.Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(createdAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateRefreshToken(userID int64, email string) (string, error) {

	createdAt := time.Now()
	expiresAt := createdAt.Add(7 * 24 * time.Hour)

	claims := JWTClaims{
		Email:     email,
		UserID:    userID,
		CreatedAt: createdAt.Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(createdAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
