package repository

import (
	"database/sql"
	"errors"
	"goapi/internal/model"
)

type VaccinationRepository struct {
	DB *sql.DB
}

func NewVaccinationRepository(db *sql.DB) *VaccinationRepository {
	return &VaccinationRepository{
		DB: db,
	}
}

func (r *VaccinationRepository) CreateVaccination(vaccination model.Vaccination) (int64, error) {
	const query = `
	INSERT INTO vaccinations (name, drug_id, dose, date)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	var id int64
	err := r.DB.QueryRow(query, vaccination.Name, vaccination.DrugID, vaccination.Dose, vaccination.Date).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *VaccinationRepository) UpdateVaccination(id int64, vaccination model.Vaccination) error {
	_, err := r.DB.Exec("UPDATE vaccinations SET name = $1, drug_id = $2, dose = $3, date = $4 WHERE id = $5",
		vaccination.Name, vaccination.DrugID, vaccination.Dose, vaccination.Date, id)
	return err
}

func (r *VaccinationRepository) GetAllVaccinations() ([]model.Vaccination, error) {
	var vaccinations []model.Vaccination
	rows, err := r.DB.Query("SELECT id, name, drug_id, dose, date FROM vaccinations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v model.Vaccination
		if err := rows.Scan(&v.ID, &v.Name, &v.DrugID, &v.Dose, &v.Date); err != nil {
			return nil, err
		}
		vaccinations = append(vaccinations, v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return vaccinations, nil
}

func (r *VaccinationRepository) DeleteVaccination(id int64) error {

	query := `DELETE FROM vaccinations WHERE id = $1;`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected, vaccination might not exist")
	}

	return nil
}
