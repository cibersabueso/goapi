package service

import (
	"goapi/internal/model"
	"goapi/internal/repository"
)

type DrugService struct {
	Repo *repository.DrugRepository
}

// CreateDrug es un método para crear un nuevo medicamento
func (s *DrugService) CreateDrug(drug model.Drug) (int64, error) {
	return s.Repo.CreateDrug(drug)
}

// UpdateDrug es un método para actualizar un medicamento existente
func (s *DrugService) UpdateDrug(id int64, drug model.Drug) error {
	// Aquí deberías implementar la lógica para actualizar el medicamento en la base de datos
	// Esto es solo un ejemplo, necesitarás implementar la lógica de actualización en tu repositorio
	return s.Repo.UpdateDrug(id, drug)
}

func (s *DrugService) GetAllDrugs() ([]model.Drug, error) {
	return s.Repo.GetAllDrugs()
}

func (s *DrugService) DeleteDrug(id int64) error {
	return s.Repo.DeleteDrug(id)
}
