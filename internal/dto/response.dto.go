package dto

type Response struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Error   string         `json:"error,omitempty"`
	Data    []UserResponse `json:"data,omitempty"`
}

type ResponseProduct struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Error   string            `json:"error,omitempty"`
	Data    []ProductResponse `json:"data,omitempty"`
}
