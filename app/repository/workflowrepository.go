package repository

import (
	"database/sql"
	"log"
	"workflowmanager/db"
	"workflowmanager/db/queries"
	"workflowmanager/models"
	"workflowmanager/repository/mapper"
)

func GetWorkflows() ([]models.Workflow, error) {
	database := db.GetDatabaseInstance()
	rows, err := database.Query(queries.GetWorkflowsQuery)

	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("error during closing the reader of the database")
		}
	}(rows)

	return mapper.WorkflowMapped(rows), nil
}
