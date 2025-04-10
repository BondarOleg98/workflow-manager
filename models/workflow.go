package models

import "time"

type Workflow struct {
	WorkflowId string
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
