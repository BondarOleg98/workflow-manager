package models

type RegisterResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
