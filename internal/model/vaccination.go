package model

import "time"

// Vaccination representa una vacunación en el sistema.
type Vaccination struct {
	ID     int64     `json:"id"`
	Name   string    `json:"name"`
	DrugID int64     `json:"drug_id"`
	Dose   int       `json:"dose"`
	Date   time.Time `json:"date"`
}
