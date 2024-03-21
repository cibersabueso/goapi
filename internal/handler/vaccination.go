package handler

import (
	"encoding/json"
	"goapi/internal/model"
	"goapi/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type VaccinationHandler struct {
	Service *service.VaccinationService
}

func NewVaccinationHandler(service *service.VaccinationService) *VaccinationHandler {
	return &VaccinationHandler{
		Service: service,
	}
}

func (h *VaccinationHandler) CreateVaccination(w http.ResponseWriter, r *http.Request) {
	var vaccination model.Vaccination
	err := json.NewDecoder(r.Body).Decode(&vaccination)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Asumiendo que vaccination.Date ya es de tipo time.Time, no es necesario parsearlo.
	// Si necesitas realizar alguna operación con la fecha, puedes hacerlo directamente.

	id, err := h.Service.CreateVaccination(vaccination)
	if err != nil {
		http.Error(w, "Error al crear la vacunación: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (h *VaccinationHandler) UpdateVaccination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid vaccination ID", http.StatusBadRequest)
		return
	}

	var vaccination model.Vaccination
	err = json.NewDecoder(r.Body).Decode(&vaccination)
	if err != nil {
		http.Error(w, "Error decoding vaccination data", http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateVaccination(id, vaccination)
	if err != nil {
		http.Error(w, "Error updating vaccination: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vaccination updated successfully"})
}

func (h *VaccinationHandler) GetAllVaccinations(w http.ResponseWriter, r *http.Request) {
	vaccinations, err := h.Service.GetAllVaccinations()
	if err != nil {
		http.Error(w, "Error al obtener las vacunaciones: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if vaccinations == nil {
		vaccinations = []model.Vaccination{} // Asegúrate de devolver un slice vacío en lugar de nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vaccinations)
}

func (h *VaccinationHandler) DeleteVaccination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid vaccination ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteVaccination(id)
	if err != nil {
		http.Error(w, "Error deleting vaccination: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vaccination deleted successfully"})
}
