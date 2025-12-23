package repository

import "workflowmanager/app/models"

type WorkflowRepository interface {
	GetWorkflowsByPagination(cursor string, pageSize int) ([]models.Workflow, error)
	GetWorkflowById(workflowId string) (models.Workflow, error)
	RemoveWorkflowById(workflowId string) (int64, error)
	SaveWorkflow(workflow models.Workflow) error
}
