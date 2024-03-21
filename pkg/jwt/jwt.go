package jwt

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Sub    int64 `json:"sub"`
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func VerifyToken(tokenString string) (*Claims, error) {

	secret := os.Getenv("JWT_SECRET_KEY")
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method in JWT")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return claims, nil
}
