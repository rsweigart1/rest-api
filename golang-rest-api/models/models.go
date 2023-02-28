package models

type AntibioticResistance struct {
	EventYearID string `json:"year"`
	State       string `json:"state"`
	Pct_Resist  string `json:"percent resistant"`
}
