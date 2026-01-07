package repository

import (
	"database/sql"
	"log"
	"log/slog"
	"workflowmanager/app/components/repository/mapper"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type PostgresWorkflowRepository struct {
	database *sql.DB
}

func NewPostgresWorkflowRepository(database *sql.DB) *PostgresWorkflowRepository {
	return &PostgresWorkflowRepository{
		database: database,
	}
}

func (postgresWorkflowRepository *PostgresWorkflowRepository) GetWorkflowsByPagination(cursor string, pageSize int) ([]models.Workflow, error) {
	database := postgresWorkflowRepository.database
	var rows *sql.Rows
	var err error

	if cursor != "" {
		rows, err = database.Query(queries.GetWorkflowsByPaginationQuery, cursor, pageSize)
	} else {
		rows, err = database.Query(queries.GetWorkflowsByPaginationWithoutCursorQuery, pageSize)
	}

	if err != nil {
		slog.Error("The error during getting all workflows from DB: %s", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("The error during closing the reader of the database")
		}
	}(rows)
	return mapper.ListEntitiesMapped([]models.Workflow{}, rows)
}

func (postgresWorkflowRepository *PostgresWorkflowRepository) GetWorkflowById(workflowId string) (models.Workflow, error) {
	database := postgresWorkflowRepository.database
	row, err := database.Query(queries.GetWorkflowByIdQuery, workflowId)
	if err != nil {
		slog.Error("The error during getting workflow by id %s from DB: %s", workflowId, err)
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
		slog.Error("The error during getting workflow by id %s from DB: %s",
			workflowId, sql.ErrNoRows)
		return models.Workflow{}, sql.ErrNoRows
	}
	return mapper.EntityMapped(models.Workflow{}, row)
}

func (postgresWorkflowRepository *PostgresWorkflowRepository) RemoveWorkflowById(workflowId string) (int64, error) {
	database := postgresWorkflowRepository.database
	resultDeletedWorkflows, err := database.Exec(queries.RemoveWorkflowByIdQuery, workflowId)
	if err != nil {
		slog.Error("The error during deleting workflows by workflow id %s from DB: %s",
			workflowId, err)
		return 0, err
	}
	rowsWorkflowsAffected, _ := resultDeletedWorkflows.RowsAffected()
	return rowsWorkflowsAffected, err
}

func (postgresWorkflowRepository *PostgresWorkflowRepository) SaveWorkflow(workflow models.Workflow) error {
	database := postgresWorkflowRepository.database
	err := database.QueryRow(queries.InsertWorkflowQuery,
		workflow.Name, workflow.CreatedAt, workflow.UpdatedAt, workflow.State).Scan(&workflow.WorkflowId)
	if err != nil {
		slog.Error("The error during saving the workflow into DB: %s", err)
		return err
	}
	for _, task := range workflow.Tasks {
		_, err = database.Exec(queries.InsertTaskQuery,
			workflow.WorkflowId.String(), task.Name, task.CreatedAt, task.UpdatedAt, workflow.State)
		if err != nil {
			slog.Error("The error during saving the task into DB: %s under workflowId - %s",
				err, workflow.WorkflowId.String())
			return err
		}
	}
	return nil
}
