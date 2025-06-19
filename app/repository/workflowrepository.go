package repository

import (
	"database/sql"
	"log"
	"workflowmanager/app/db"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
	"workflowmanager/app/repository/mapper"
)

func GetWorkflows() ([]models.Workflow, error) {
	database := db.GetDatabaseInstance()
	rows, err := database.Query(queries.GetWorkflowsQuery)

	if err != nil {
		log.Fatalf("The error during getting all worflows from DB: %s", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("The error during closing the reader of the database")
		}
	}(rows)
	log.Printf("Worflows were retrieved from DB")
	return mapper.WorkflowMapped(rows), nil
}
