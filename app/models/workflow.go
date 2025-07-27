package models

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Workflow struct {
	WorkflowId ulid.ULID `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
