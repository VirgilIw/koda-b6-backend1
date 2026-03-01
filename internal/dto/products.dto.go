package dto

type ProductResponse struct {
	Id          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Rating      float64 `json:"rating,omitempty"`
	Stock       int     `json:"stock,omitempty"`
	Images      string  `json:"images,omitempty"`
}

type ProductRequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Stock       int     `json:"stock"`
	Images      string  `json:"images"`
}
