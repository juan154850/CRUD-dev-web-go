package models

type Movie struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	Director     string  `json:"director"`
	Genre        string  `json:"genre"`
	Description  string  `json:"description"`
	Year         int     `json:"year"`
	Calification float32 `json:"calification"`
	Duration     int     `json:"duration"`
}