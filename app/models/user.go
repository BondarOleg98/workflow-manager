package models

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type User struct {
	Id        ulid.ULID
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	LastLogin *time.Time
}
