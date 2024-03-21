package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken genera un nuevo token JWT para un usuario dado.
func GenerateToken(userID int64) (string, error) {
	// Obtener la duración del token de una variable de entorno o usar un valor predeterminado.
	tokenDurationStr := os.Getenv("JWT_EXPIRATION")
	if tokenDurationStr == "" {
		tokenDurationStr = "24h" // Valor predeterminado de 24 horas si no se establece la variable de entorno
	}

	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		return "", err
	}

	// Crear el token con los claims (reclamaciones) necesarios.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(tokenDuration).Unix(),
	})

	// Firmar el token con una clave secreta definida en una variable de entorno.
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
