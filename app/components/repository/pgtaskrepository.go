package repository

import (
	"database/sql"
	"log"
	"log/slog"
	"workflowmanager/app/components/repository/mapper"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type PostgresTaskRepository struct {
	database *sql.DB
}

func NewPostgresTaskRepository(database *sql.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{
		database: database,
	}
}

func (postgresTaskRepository *PostgresTaskRepository) RemoveTaskById(taskId string) (int64, error) {
	resultDeletedTask, err := postgresTaskRepository.database.Exec(queries.RemoveTaskByIdQuery, taskId)
	if err != nil {
		slog.Error("The error during deleting the task by Id %s from DB: %s",
			taskId, err)
		return 0, err
	}
	rowsTasksAffected, _ := resultDeletedTask.RowsAffected()
	slog.Info("Task by taskId - %s was removed %d", taskId, rowsTasksAffected)
	return rowsTasksAffected, err
}

func (postgresTaskRepository *PostgresTaskRepository) GetTasksByPagination(cursor string, pageSize int) ([]models.Task, error) {
	var rows *sql.Rows
	var err error
	if cursor != "" {
		rows, err = postgresTaskRepository.database.
			Query(queries.GetTasksByPaginationQuery, cursor, pageSize)
	} else {
		rows, err = postgresTaskRepository.database.
			Query(queries.GetTasksByPaginationWithoutCursorQuery, pageSize)
	}
	if err != nil {
		slog.Error("The error during getting tasks from DB: %s", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("The error during closing the reader of the database")
		}
	}(rows)
	return mapper.ListEntitiesMapped([]models.Task{}, rows)
}

func (postgresTaskRepository *PostgresTaskRepository) GetTaskById(workflowId string) (models.Task, error) {
	row, err := postgresTaskRepository.database.Query(queries.GetTaskByIdQuery, workflowId)
	if err != nil {
		slog.Info("The error during getting task by id %s from DB: %s", workflowId, err)
		return models.Task{}, err
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal("The error during closing the reader of the database")
		}
	}(row)
	isSqlResultEmpty := row.Next()
	if !isSqlResultEmpty {
		slog.Error("The error during getting task by id %s from DB: %s",
			workflowId, sql.ErrNoRows)
		return models.Task{}, sql.ErrNoRows
	}
	return mapper.EntityMapped(models.Task{}, row)
}

func (postgresTaskRepository *PostgresTaskRepository) GetTasksByWorkflowId(workflowId string) ([]models.Task, error) {
	rows, err := postgresTaskRepository.database.Query(queries.GetTasksByWorkflowIdQuery, workflowId)
	if err != nil {
		slog.Error("The error during getting tasks by workflowId %s from DB: %s", workflowId, err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("The error during closing the reader of the database")
		}
	}(rows)
	return mapper.ListEntitiesMapped([]models.Task{}, rows)
}
