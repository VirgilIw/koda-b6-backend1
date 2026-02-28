package model

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Stock       int     `json:"stock"`
	Images      string  `json:"images"`
}
