package service

import (
	"goapi/internal/model"
	"goapi/internal/repository"
)

type DrugService struct {
	Repo *repository.DrugRepository
}

func (s *DrugService) CreateDrug(drug model.Drug) (int64, error) {
	return s.Repo.CreateDrug(drug)
}

func (s *DrugService) UpdateDrug(id int64, drug model.Drug) error {

	return s.Repo.UpdateDrug(id, drug)
}

func (s *DrugService) GetAllDrugs() ([]model.Drug, error) {
	return s.Repo.GetAllDrugs()
}

func (s *DrugService) DeleteDrug(id int64) error {
	return s.Repo.DeleteDrug(id)
}
