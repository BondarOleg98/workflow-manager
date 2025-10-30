package models

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type RefreshToken struct {
	Id        ulid.ULID
	UserId    ulid.ULID
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
	Revoked   bool
}
