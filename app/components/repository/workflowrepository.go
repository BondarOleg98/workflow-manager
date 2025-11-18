package repository

import (
	"database/sql"
	"log"
	"workflowmanager/app/components/repository/mapper"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type WorkflowRepository struct {
	database *sql.DB
}

func NewWorkflowRepository(database *sql.DB) *WorkflowRepository {
	return &WorkflowRepository{
		database: database,
	}
}

func (workflowRepository *WorkflowRepository) GetWorkflowsByPagination(cursor string, pageSize int) ([]models.Workflow, error) {
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
	return mapper.ListEntitiesMapped([]models.Workflow{}, rows)
}

func (workflowRepository *WorkflowRepository) GetWorkflowById(workflowId string) (models.Workflow, error) {
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
	return mapper.EntityMapped(models.Workflow{}, row)
}

func (workflowRepository *WorkflowRepository) RemoveWorkflowById(workflowId string) (int64, error) {
	database := workflowRepository.database
	resultDeletedWorkflows, err := database.Exec(queries.RemoveWorkflowByIdQuery, workflowId)
	if err != nil {
		log.Printf("The error during deleting workflows by workflow id %s from DB: %s",
			workflowId, err)
		return 0, err
	}
	rowsWorkflowsAffected, _ := resultDeletedWorkflows.RowsAffected()
	return rowsWorkflowsAffected, err
}

func (workflowRepository *WorkflowRepository) SaveWorkflow(workflow models.Workflow) error {
	database := workflowRepository.database
	for _, task := range workflow.Tasks {
		_, err := database.Exec(queries.InsertWorkflowQuery,
			workflow.WorkflowId.String(), workflow.Name, workflow.CreatedAt, workflow.UpdatedAt, workflow.State)
		if err != nil {
			log.Printf("The error during saving the workflow into DB: %s", err)
			return err
		}
		_, err = database.Exec(queries.InsertTaskQuery,
			task.TaskId.String(), workflow.WorkflowId.String(), task.Name, task.CreatedAt, task.UpdatedAt, workflow.State)
		if err != nil {
			log.Printf("The error during saving the task into DB: %s under workflowId - %s",
				err, workflow.WorkflowId.String())
			return err
		}
	}
	return nil
}
