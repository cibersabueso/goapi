package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int64) (string, error) {

	tokenDurationStr := os.Getenv("JWT_EXPIRATION")
	if tokenDurationStr == "" {
		tokenDurationStr = "24h"
	}

	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(tokenDuration).Unix(),
	})

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("la clave secreta del JWT no está definida")
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
