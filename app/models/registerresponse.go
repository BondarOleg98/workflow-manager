package models

type RegisterResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
