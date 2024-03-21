package handler

import (
	"database/sql"
	"encoding/json"
	"goapi/internal/auth"
	"goapi/internal/model"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Signup maneja el registro de usuarios
func Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("Error al decodificar las credenciales:", err)
			http.Error(w, "Error al decodificar las credenciales", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error al hashear la contraseña:", err)
			http.Error(w, "Error al hashear la contraseña", http.StatusInternalServerError)
			return
		}

		stmt := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
		var userId int
		err = db.QueryRow(stmt, user.Name, user.Email, hashedPassword).Scan(&userId)
		if err != nil {
			log.Println("Error al registrar el usuario:", err)
			http.Error(w, "Error al registrar el usuario", http.StatusInternalServerError)
			return
		}

		userId64 := int64(userId)
		token, err := auth.GenerateToken(userId64)
		if err != nil {
			log.Println("Error al generar el token:", err)
			http.Error(w, "Error al generar el token", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{"token": token}); err != nil {
			log.Println("Error al codificar la respuesta JSON:", err)
			http.Error(w, "Error al codificar la respuesta JSON", http.StatusInternalServerError)
		}
	}
}

// Login maneja el inicio de sesión de usuarios
func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			http.Error(w, "Error al decodificar las credenciales", http.StatusBadRequest)
			return
		}

		var u model.User
		stmt := `SELECT id, password FROM users WHERE email = $1`
		err = db.QueryRow(stmt, credentials.Email).Scan(&u.ID, &u.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "El usuario no existe", http.StatusUnauthorized)
			} else {
				log.Println("Error al consultar la base de datos:", err)
				http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
			}
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(credentials.Password))
		if err != nil {
			http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
			return
		}

		token, err := auth.GenerateToken(u.ID)
		if err != nil {
			log.Println("Error al generar el token:", err)
			http.Error(w, "Error al generar el token", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{"token": token}); err != nil {
			log.Println("Error al codificar la respuesta JSON:", err)
			http.Error(w, "Error al codificar la respuesta JSON", http.StatusInternalServerError)
		}
	}
}
