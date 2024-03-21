package handler

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// SetupRoutes configura las rutas de la aplicación.
func SetupRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/signup", Signup(db)).Methods("POST")
	r.HandleFunc("/login", Login(db)).Methods("POST")
	// Configurar las demás rutas aquí
}
