package services

import (
	"github.com/oklog/ulid/v2"
	"log"
	"time"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/models"
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
		log.Printf("Workflows were retrieved")
	}
	return
}

func (workflowService *WorkflowService) GetWorkflowById(workflowId string) (workflow models.Workflow, err error) {
	workflow, err = workflowService.workflowRepository.GetWorkflowById(workflowId)
	if err == nil {
		log.Printf("Workflow by id - %s was retrieved", workflowId)
	}
	return
}

func (workflowService *WorkflowService) RemoveWorkflowById(workflowId string) error {
	rowsWorkflowsAffected, err := workflowService.workflowRepository.RemoveWorkflowById(workflowId)
	if err == nil {
		log.Printf("Workflow by id - %s was removed %d", workflowId, rowsWorkflowsAffected)
		return nil
	}
	return err
}

func (workflowService *WorkflowService) SaveWorkflow(workflow models.Workflow) error {
	createdAt := time.Now()
	workflowId := ulid.Make()
	workflow.WorkflowId = workflowId
	workflow.State = models.CREATED
	for indexTask := range workflow.Tasks {
		taskId := ulid.Make()
		workflow.Tasks[indexTask].TaskId = taskId
		workflow.Tasks[indexTask].CreatedAt = createdAt
		workflow.Tasks[indexTask].UpdatedAt = createdAt
		workflow.Tasks[indexTask].State = models.CREATED
	}
	err := workflowService.workflowRepository.SaveWorkflow(workflow)
	if err != nil {
		return err
	}
	log.Printf("The workflow with id: %s was created", workflowId)
	return nil
}
