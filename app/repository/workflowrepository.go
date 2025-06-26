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
		log.Printf("The error during getting all workflows from DB: %s", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("The error during closing the reader of the database")
		}
	}(rows)

	return mapper.WorkflowsListMapped(rows)
}

func GetWorkflowById(workflowId string) (models.Workflow, error) {
	database := db.GetDatabaseInstance()
	row, err := database.Query(queries.GetWorkflowByIdQuery, workflowId)
	if err != nil {
		log.Printf("The error during getting workflow by id %s from DB: %s", workflowId, err)
		return models.Workflow{}, err
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal("The error during closing the reader of the database")
		}
	}(row)
	isSqlResultEmpty := row.Next()
	if !isSqlResultEmpty {
		log.Printf("The error during getting workflow by id %s from DB: %s",
			workflowId, sql.ErrNoRows)
		return models.Workflow{}, sql.ErrNoRows
	}
	return mapper.WorkflowMapped(row)
}
