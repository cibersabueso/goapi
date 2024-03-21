package service

import (
	"errors"
	"goapi/internal/model"
	"goapi/internal/repository"
	"time"
)

type VaccinationService struct {
	Repo     *repository.VaccinationRepository
	DrugRepo *repository.DrugRepository // Asegúrate de tener esto para acceder al repositorio de Drug
}

func NewVaccinationService(repo *repository.VaccinationRepository, drugRepo *repository.DrugRepository) *VaccinationService {
	return &VaccinationService{
		Repo:     repo,
		DrugRepo: drugRepo,
	}
}

func (s *VaccinationService) CreateVaccination(vaccination model.Vaccination) (int64, error) {
	// Primero, obtener la información del medicamento (drug) para validar la dosis y la fecha
	drug, err := s.DrugRepo.GetDrugByID(vaccination.DrugID)
	if err != nil {
		return 0, err // Manejar el error adecuadamente
	}

	// Validar la dosis
	if vaccination.Dose < drug.MinDose || vaccination.Dose > drug.MaxDose {
		return 0, errors.New("la dosis está fuera del rango permitido")
	}

	// No es necesario convertir vaccination.Date, ya es de tipo time.Time

	// Convertir drug.AvailableAt de string a time.Time
	availableAtTime, err := time.Parse(time.RFC3339, drug.AvailableAt) // Usar el formato RFC3339 que es compatible con ISO 8601
	if err != nil {
		return 0, err
	}

	// Validar la fecha
	if vaccination.Date.Before(availableAtTime) {
		return 0, errors.New("la fecha de vacunación es anterior a la fecha de disponibilidad del medicamento")
	}

	// Si todo es válido, proceder a crear la vacunación
	return s.Repo.CreateVaccination(vaccination)
}

func (s *VaccinationService) UpdateVaccination(id int64, vaccination model.Vaccination) error {
	// Aquí puedes agregar validaciones adicionales si es necesario
	return s.Repo.UpdateVaccination(id, vaccination)
}

func (s *VaccinationService) GetAllVaccinations() ([]model.Vaccination, error) {
	return s.Repo.GetAllVaccinations()
}

func (s *VaccinationService) DeleteVaccination(id int64) error {
	// Implementa la lógica para eliminar la vacunación por ID
	return s.Repo.DeleteVaccination(id)
}
