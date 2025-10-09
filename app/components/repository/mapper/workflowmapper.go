package mapper

import (
	"database/sql"
	"log"
	"workflowmanager/app/models"
)

func WorkflowsListMapped(rows *sql.Rows) ([]models.Workflow, error) {
	var workflows []models.Workflow
	for rows.Next() {
		workflow := models.Workflow{}
		if err := rows.Scan(&workflow.WorkflowId, &workflow.Name, &workflow.CreatedAt, &workflow.UpdatedAt); err != nil {
			log.Fatalf("The error during mapping data from DB %s", err)
			return nil, err
		}
		workflows = append(workflows, workflow)
	}
	return workflows, nil
}

func WorkflowMapped(row *sql.Rows) (models.Workflow, error) {
	var err error
	workflow := models.Workflow{}
	if err = row.Scan(&workflow.WorkflowId, &workflow.Name, &workflow.CreatedAt, &workflow.UpdatedAt); err != nil {
		log.Printf("The error during mapping data from DB %s", err)
		return workflow, err
	}
	return workflow, err
}
