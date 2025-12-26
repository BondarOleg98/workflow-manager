package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	LastLogin *time.Time
}
