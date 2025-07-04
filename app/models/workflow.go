package models

import "time"

type Workflow struct {
	WorkflowId string `json:"id"`
	Name       string `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
