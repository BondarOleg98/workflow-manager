package models

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Workflow struct {
	WorkflowId ulid.ULID `json:"workflow_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	State      State     `json:"state"`
	Tasks      []Task    `json:"tasks"`
}
