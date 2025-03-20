package repository

import (
	"database/sql"
	"workflowmanager/db"
	"workflowmanager/db/queries"
)

// GetWorkflows TODO: add handling of the error
func GetWorkflows() *sql.Rows {
	pool := db.Pool{
		Username:     "postgres",
		DatabaseName: "postgres",
	}
	database := db.OpenDatabaseConnection(pool)
	rows, _ := database.Query(queries.GetWorkflowsQuery)
	return rows
}
