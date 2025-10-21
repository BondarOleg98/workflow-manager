package models

import (
	"errors"
	"github.com/oklog/ulid/v2"
	"time"
)

type contextKey string

const UserIdKey contextKey = "userId"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("token has expired")
	ErrEmailInUse         = errors.New("email already in use")
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type RefreshToken struct {
	Id        ulid.ULID
	UserId    ulid.ULID
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
	Revoked   bool
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
