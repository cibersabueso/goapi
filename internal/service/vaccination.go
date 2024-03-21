package service

import (
	"errors"
	"goapi/internal/model"
	"goapi/internal/repository"
	"time"
)

type VaccinationService struct {
	Repo     *repository.VaccinationRepository
	DrugRepo *repository.DrugRepository
}

func NewVaccinationService(repo *repository.VaccinationRepository, drugRepo *repository.DrugRepository) *VaccinationService {
	return &VaccinationService{
		Repo:     repo,
		DrugRepo: drugRepo,
	}
}

func (s *VaccinationService) CreateVaccination(vaccination model.Vaccination) (int64, error) {

	drug, err := s.DrugRepo.GetDrugByID(vaccination.DrugID)
	if err != nil {
		return 0, err
	}

	if vaccination.Dose < drug.MinDose || vaccination.Dose > drug.MaxDose {
		return 0, errors.New("la dosis está fuera del rango permitido")
	}

	availableAtTime, err := time.Parse(time.RFC3339, drug.AvailableAt)
	if err != nil {
		return 0, err
	}

	if vaccination.Date.Before(availableAtTime) {
		return 0, errors.New("la fecha de vacunación es anterior a la fecha de disponibilidad del medicamento")
	}

	return s.Repo.CreateVaccination(vaccination)
}

func (s *VaccinationService) UpdateVaccination(id int64, vaccination model.Vaccination) error {

	return s.Repo.UpdateVaccination(id, vaccination)
}

func (s *VaccinationService) GetAllVaccinations() ([]model.Vaccination, error) {
	return s.Repo.GetAllVaccinations()
}

func (s *VaccinationService) DeleteVaccination(id int64) error {

	return s.Repo.DeleteVaccination(id)
}
