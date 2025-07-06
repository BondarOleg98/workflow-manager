package models

import (
	"github.com/google/uuid"
	"time"
)

type Workflow struct {
	WorkflowId uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
