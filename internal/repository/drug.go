package repository

import (
	"database/sql"
	"goapi/internal/model"
)

type DrugRepository struct {
	DB *sql.DB
}

// CreateDrug es un método para crear un nuevo medicamento en la base de datos
func (r *DrugRepository) CreateDrug(drug model.Drug) (int64, error) {
	stmt, err := r.DB.Prepare("INSERT INTO drugs(name, approved, min_dose, max_dose, available_at) VALUES($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var drugID int64
	err = stmt.QueryRow(drug.Name, drug.Approved, drug.MinDose, drug.MaxDose, drug.AvailableAt).Scan(&drugID)
	if err != nil {
		return 0, err
	}

	return drugID, nil
}

// UpdateDrug es un método para actualizar un medicamento existente en la base de datos
func (r *DrugRepository) UpdateDrug(id int64, drug model.Drug) error {
	stmt, err := r.DB.Prepare("UPDATE drugs SET name = $1, approved = $2, min_dose = $3, max_dose = $4, available_at = $5 WHERE id = $6")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(drug.Name, drug.Approved, drug.MinDose, drug.MaxDose, drug.AvailableAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *DrugRepository) GetAllDrugs() ([]model.Drug, error) {
	rows, err := r.DB.Query("SELECT id, name, approved, min_dose, max_dose, available_at FROM drugs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drugs []model.Drug
	for rows.Next() {
		var d model.Drug
		if err := rows.Scan(&d.ID, &d.Name, &d.Approved, &d.MinDose, &d.MaxDose, &d.AvailableAt); err != nil {
			return nil, err
		}
		drugs = append(drugs, d)
	}

	return drugs, nil
}

func (r *DrugRepository) DeleteDrug(id int64) error {
	stmt, err := r.DB.Prepare("DELETE FROM drugs WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// GetDrugByID busca un medicamento por su ID y lo devuelve.
func (r *DrugRepository) GetDrugByID(id int64) (*model.Drug, error) {
	var drug model.Drug
	query := `SELECT id, name, approved, min_dose, max_dose, available_at FROM drugs WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&drug.ID, &drug.Name, &drug.Approved, &drug.MinDose, &drug.MaxDose, &drug.AvailableAt)
	if err != nil {
		return nil, err
	}
	return &drug, nil
}
