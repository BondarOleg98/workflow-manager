package repository

import (
	"database/sql"
	"log"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
	"workflowmanager/app/repository/mapper"
)

type WorkflowRepository struct {
	database *sql.DB
}

func InitWorkflowRepository(database *sql.DB) WorkflowRepository {
	return WorkflowRepository{
		database: database,
	}
}

func (workflowRepository WorkflowRepository) GetWorkflowsByPagination(cursor string, pageSize int) ([]models.Workflow, error) {
	database := workflowRepository.database
	var rows *sql.Rows
	var err error

	if cursor != "" {
		rows, err = database.Query(queries.GetWorkflowsByPaginationQuery, cursor, pageSize)
	} else {
		rows, err = database.Query(queries.GetWorkflowsByPaginationWithoutCursorQuery, pageSize)
	}

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

func (workflowRepository WorkflowRepository) GetWorkflowById(workflowId string) (models.Workflow, error) {
	database := workflowRepository.database
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

func (workflowRepository WorkflowRepository) removeTasksByWorkflowId(workflowId string) (int64, error) {
	database := workflowRepository.database
	_, err := workflowRepository.removeActionsByWorkflowId(workflowId)
	if err != nil {
		return 0, err
	}
	resultDeletedTasks, err := database.Exec(queries.RemoveTasksByWorkflowIdQuery, workflowId)
	if err != nil {
		log.Printf("The error during deleting tasks by workflow id %s from DB: %s",
			workflowId, err)
		return 0, err
	}
	rowsTasksAffected, _ := resultDeletedTasks.RowsAffected()
	log.Printf("Tasks by workflowId - %s were removed %d", workflowId, rowsTasksAffected)
	return rowsTasksAffected, err
}

func (workflowRepository WorkflowRepository) removeActionsByWorkflowId(workflowId string) (int64, error) {
	database := workflowRepository.database
	resultDeletedActions, err := database.Exec(queries.RemoveActionsByTaskIdQuery, workflowId)
	if err != nil {
		log.Printf("The error during deleting actions by workflow id %s from DB: %s",
			workflowId, err)
		return 0, err
	}
	rowsActionsAffected, _ := resultDeletedActions.RowsAffected()
	log.Printf("Actions by workflowId - %s were removed %d", workflowId, rowsActionsAffected)
	return rowsActionsAffected, err
}

func (workflowRepository WorkflowRepository) RemoveWorkflowById(workflowId string) (int64, error) {
	database := workflowRepository.database
	_, err := workflowRepository.removeTasksByWorkflowId(workflowId)
	if err != nil {
		return 0, err
	}
	resultDeletedWorkflows, err := database.Exec(queries.RemoveWorkflowByIdQuery, workflowId)
	if err != nil {
		log.Printf("The error during deleting workflows by workflow id %s from DB: %s",
			workflowId, err)
		return 0, err
	}
	rowsWorkflowsAffected, _ := resultDeletedWorkflows.RowsAffected()
	return rowsWorkflowsAffected, err
}

func (workflowRepository WorkflowRepository) SaveWorkflow(workflow models.Workflow) error {
	database := workflowRepository.database
	_, err := database.Exec(queries.InsertWorkflowQuery,
		workflow.WorkflowId.String(), workflow.Name, workflow.UpdatedAt, workflow.CreatedAt)
	if err != nil {
		log.Printf("The error during saving the workflow into DB: %s", err)
		return err
	}
	return nil
}
