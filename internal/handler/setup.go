package handler

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/signup", Signup(db)).Methods("POST")
	r.HandleFunc("/login", Login(db)).Methods("POST")

}
