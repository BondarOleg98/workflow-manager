package mapper

import (
	"database/sql"
	"log"
	"time"
	"workflowmanager/app/models"
)

func WorkflowMapped(rows *sql.Rows) []models.Workflow {
	var workflows []models.Workflow
	for rows.Next() {
		var (
			workflowId string
			name       string
			createdAt  time.Time
			updatedAt  time.Time
		)
		if err := rows.Scan(&workflowId, &name, &createdAt, &updatedAt); err != nil {
			log.Fatalf("The error during mapping data from DB %s", err)
		}
		workflows = append(workflows, models.Workflow{
			WorkflowId: workflowId,
			Name:       name,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		})
	}
	return workflows
}
