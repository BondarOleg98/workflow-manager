package models

import (
	"github.com/google/uuid"
	"time"
)

type Workflow struct {
	WorkflowId uuid.UUID `json:"workflow_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	State      State     `json:"state"`
	Tasks      []Task    `json:"tasks" mapper:"omit"`
}
