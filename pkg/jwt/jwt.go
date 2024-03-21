package jwt

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// Claims estructura para los claims del JWT que incluye el UserID.
type Claims struct {
	Sub    int64 `json:"sub"` // Asegúrate de que el tipo de dato aquí sea el esperado por tu aplicación.
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// VerifyToken verifica el token JWT y devuelve los claims si el token es válido.
func VerifyToken(tokenString string) (*Claims, error) {
	// Aquí se asume que tu secreto está almacenado en la variable de entorno JWT_SECRET_KEY
	secret := os.Getenv("JWT_SECRET_KEY")
	claims := &Claims{} // Instancia de Claims personalizados

	// Parsear el token con los Claims personalizados
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

	return claims, nil // Devolver los Claims personalizados
}
