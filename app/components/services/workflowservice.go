package services

import (
	"log/slog"
	"time"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/models"
	"workflowmanager/app/models/filtration"
)

type WorkflowService struct {
	workflowRepository repository.WorkflowRepository
}

func NewWorkflowService(workflowRepository repository.WorkflowRepository) *WorkflowService {
	return &WorkflowService{
		workflowRepository: workflowRepository,
	}
}

func (workflowService *WorkflowService) GetWorkflowsByPagination(cursor string, pageSize int) (workflows []models.Workflow, err error) {
	workflows, err = workflowService.workflowRepository.GetWorkflowsByPagination(cursor, pageSize)
	if err == nil {
		slog.Info("Workflows were retrieved")
	}
	return
}

func (workflowService *WorkflowService) GetWorkflowById(workflowId string) (workflow models.Workflow, err error) {
	workflow, err = workflowService.workflowRepository.GetWorkflowById(workflowId)
	if err == nil {
		slog.Info("Workflow was retrieved", "id", workflowId)
	}
	return
}

func (workflowService *WorkflowService) RemoveWorkflowById(workflowId string) error {
	rowsWorkflowsAffected, err := workflowService.workflowRepository.RemoveWorkflowById(workflowId)
	if err == nil {
		slog.Info("Workflow by id was removed:", "workflowId", workflowId, "affectedRows", rowsWorkflowsAffected)
		return nil
	}
	return err
}

func (workflowService *WorkflowService) SaveWorkflow(workflow models.Workflow) (*models.Workflow, error) {
	createdAt := time.Now()
	workflow.State = models.CREATED
	workflow.CreatedAt = createdAt
	workflow.UpdatedAt = createdAt
	for indexTask := range workflow.Tasks {
		workflow.Tasks[indexTask].CreatedAt = createdAt
		workflow.Tasks[indexTask].UpdatedAt = createdAt
		workflow.Tasks[indexTask].State = models.CREATED
	}
	createdWorkflow, err := workflowService.workflowRepository.SaveWorkflow(workflow)
	if err != nil {
		return createdWorkflow, err
	}
	slog.Info("The workflow was created with", "name", createdWorkflow.Name)
	return createdWorkflow, nil
}

func (workflowService *WorkflowService) GetWorkflowsByFiltration(filtration filtration.Filtration) (workflows []models.Workflow, err error) {
	workflows, err = workflowService.workflowRepository.GetWorkflowsByFiltration(filtration)
	if err == nil {
		slog.Info("Workflows after filtering were retrieved")
	}
	return
}
