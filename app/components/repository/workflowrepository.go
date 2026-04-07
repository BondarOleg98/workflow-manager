package repository

import (
	"workflowmanager/app/models"
	"workflowmanager/app/models/filtration"
)

type WorkflowRepository interface {
	GetWorkflowsByPagination(cursor string, pageSize int) ([]models.Workflow, error)
	GetWorkflowById(workflowId string) (models.Workflow, error)
	RemoveWorkflowById(workflowId string) (int64, error)
	SaveWorkflow(workflow models.Workflow) error
	GetWorkflowsByFiltration(filtration filtration.Filtration) ([]models.Workflow, error)
}
