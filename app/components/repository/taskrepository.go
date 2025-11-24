package repository

import "workflowmanager/app/models"

type TaskRepository interface {
	RemoveTaskById(taskId string) (int64, error)
	GetTasksByPagination(cursor string, pageSize int) ([]models.Task, error)
	GetTaskById(workflowId string) (models.Task, error)
	GetTasksByWorkflowId(workflowId string) ([]models.Task, error)
}
