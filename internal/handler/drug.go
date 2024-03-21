package handler

import (
	"encoding/json"
	"goapi/internal/model"
	"goapi/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // Asegúrate de tener esta dependencia importada para manejar las variables de ruta
)

type DrugHandler struct {
	Service *service.DrugService
}

// CreateDrug es un método para crear una nueva "drug"
func (h *DrugHandler) CreateDrug(w http.ResponseWriter, r *http.Request) {
	var drug model.Drug
	if err := json.NewDecoder(r.Body).Decode(&drug); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	drugID, err := h.Service.CreateDrug(drug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": drugID})
}

// UpdateDrug es un método para actualizar una "drug" existente
func (h *DrugHandler) UpdateDrug(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la URL
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid drug ID", http.StatusBadRequest)
		return
	}

	// Decodificar el cuerpo de la solicitud en una instancia de Drug
	var drug model.Drug
	if err := json.NewDecoder(r.Body).Decode(&drug); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Llamar al servicio para actualizar la "drug"
	err = h.Service.UpdateDrug(id, drug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar respuesta de éxito
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Drug updated successfully"})
}

// GetAllDrugs es un método para obtener todas las "drugs"
func (h *DrugHandler) GetAllDrugs(w http.ResponseWriter, r *http.Request) {
	drugs, err := h.Service.GetAllDrugs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drugs)
}

func (h *DrugHandler) DeleteDrug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid drug ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteDrug(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Drug deleted successfully"})
}
