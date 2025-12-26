package models

import (
	"github.com/google/uuid"
	"time"
)

type RefreshToken struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
	Revoked   bool
}
