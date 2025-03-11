package models

import "time"

type Workflow struct {
	name       string
	workflowId string
	createdAt  time.Time
	updatedAt  time.Time
}
