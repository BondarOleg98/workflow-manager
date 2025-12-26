package models

import (
	"github.com/google/uuid"
	"time"
)

type Action struct {
	ActionId  uuid.UUID `json:"action_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	State     State     `json:"state"`
}
