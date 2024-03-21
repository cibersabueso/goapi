package model

type Drug struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Approved    bool   `json:"approved"`
	MinDose     int    `json:"min_dose"`
	MaxDose     int    `json:"max_dose"`
	AvailableAt string `json:"available_at"`
}
