package repository

import (
	"database/sql"
	"log"
	"workflowmanager/app/db/queries"
)

type PostgresTaskRepository struct {
	database *sql.DB
}

func NewTaskRepository(database *sql.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{
		database: database,
	}
}

func (postgresTaskRepository *PostgresTaskRepository) RemoveTaskById(taskId string) (int64, error) {
	database := postgresTaskRepository.database
	resultDeletedTask, err := database.Exec(queries.RemoveTaskByIdQuery, taskId)
	if err != nil {
		log.Printf("The error during deleting the task by Id %s from DB: %s",
			taskId, err)
		return 0, err
	}
	rowsTasksAffected, _ := resultDeletedTask.RowsAffected()
	log.Printf("Task by taskId - %s was removed %d", taskId, rowsTasksAffected)
	return rowsTasksAffected, err
}
