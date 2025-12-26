package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	TaskId    uuid.UUID `json:"task_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	State     State     `json:"state"`
	Action    Action    `json:"-" mapper:"omit"`
}
