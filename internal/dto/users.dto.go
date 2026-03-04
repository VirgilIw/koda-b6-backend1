package dto

type UserResponse struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UpdateUserRequest struct {
	Email    string
	Password string
}
