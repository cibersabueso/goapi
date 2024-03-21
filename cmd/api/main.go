package main

import (
	"goapi/internal/config"
	"goapi/internal/handler"
	"goapi/internal/middleware"
	"goapi/internal/repository"
	"goapi/internal/service"
	"goapi/pkg/postgres"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se encontró el archivo .env o no se pudo cargar")
	} else {
		log.Println("Variables de entorno cargadas correctamente")
		log.Printf("JWT_SECRET_KEY es: %s", os.Getenv("JWT_SECRET_KEY"))
		log.Printf("JWT_EXPIRATION es: %s", os.Getenv("JWT_EXPIRATION"))
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	db, err := postgres.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/signup", handler.Signup(db)).Methods("POST")
	r.HandleFunc("/login", handler.Login(db)).Methods("POST")

	drugHandler := &handler.DrugHandler{
		Service: &service.DrugService{
			Repo: &repository.DrugRepository{
				DB: db,
			},
		},
	}

	// Ruta para crear un nuevo medicamento con validación de JWT
	r.Handle("/drugs", middleware.ValidateJWT(http.HandlerFunc(drugHandler.CreateDrug))).Methods("POST")
	r.Handle("/drugs/{id}", middleware.ValidateJWT(http.HandlerFunc(drugHandler.UpdateDrug))).Methods("PUT")
	r.Handle("/drugs", middleware.ValidateJWT(http.HandlerFunc(drugHandler.GetAllDrugs))).Methods("GET")
	r.Handle("/drugs/{id}", middleware.ValidateJWT(http.HandlerFunc(drugHandler.DeleteDrug))).Methods("DELETE")

	// Crear instancias para manejo de vacunaciones
	vaccinationRepo := repository.NewVaccinationRepository(db)
	vaccinationService := service.NewVaccinationService(vaccinationRepo, drugHandler.Service.Repo)
	vaccinationHandler := handler.NewVaccinationHandler(vaccinationService)

	// Ruta para crear una nueva vacunación con validación de JWT
	r.Handle("/vaccination", middleware.ValidateJWT(http.HandlerFunc(vaccinationHandler.CreateVaccination))).Methods("POST")
	r.Handle("/vaccination/{id}", middleware.ValidateJWT(http.HandlerFunc(vaccinationHandler.UpdateVaccination))).Methods("PUT")
	r.Handle("/vaccination", middleware.ValidateJWT(http.HandlerFunc(vaccinationHandler.GetAllVaccinations))).Methods("GET")
	r.Handle("/vaccination/{id}", middleware.ValidateJWT(http.HandlerFunc(vaccinationHandler.DeleteVaccination))).Methods("DELETE")

	// Iniciar el servidor
	log.Printf("Servidor escuchando en %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
