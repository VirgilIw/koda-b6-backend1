package dto

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UpdateUserRequest struct {
	Email    string
	Password string
}
