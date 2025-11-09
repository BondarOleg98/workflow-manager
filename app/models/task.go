package models

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Task struct {
	TaskId    ulid.ULID `json:"task_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	State     State     `json:"state"`
	Action    Action    `json:"action"`
}
