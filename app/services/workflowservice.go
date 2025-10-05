package services

import (
	"github.com/oklog/ulid/v2"
	"log"
	"time"
	"workflowmanager/app/db"
	"workflowmanager/app/models"
	"workflowmanager/app/repository"
)

type WorkflowService struct {
	workflowRepository repository.WorkflowRepository
}

func InitWorkflowService() WorkflowService {
	return WorkflowService{
		workflowRepository: repository.InitWorkflowRepository(db.GetDatabaseInstance()),
	}
}

func (workflowService WorkflowService) GetWorkflowsByPagination(cursor string, pageSize int) (workflows []models.Workflow, err error) {
	workflows, err = workflowService.workflowRepository.GetWorkflowsByPagination(cursor, pageSize)
	if err == nil {
		log.Printf("Workflows were retrieved")
	}
	return
}

func (workflowService WorkflowService) GetWorkflowById(workflowId string) (workflow models.Workflow, err error) {
	workflow, err = workflowService.workflowRepository.GetWorkflowById(workflowId)
	if err == nil {
		log.Printf("Workflow by id - %s was retrieved", workflowId)
	}
	return
}

func (workflowService WorkflowService) RemoveWorkflowById(workflowId string) error {
	rowsWorkflowsAffected, err := workflowService.workflowRepository.RemoveWorkflowById(workflowId)
	if err == nil {
		log.Printf("Workflow by id - %s was removed %d", workflowId, rowsWorkflowsAffected)
		return nil
	}
	return err
}

func (workflowService WorkflowService) SaveWorkflow(workflow models.Workflow) error {
	workflow.CreatedAt = time.Now()
	workflow.UpdatedAt = workflow.CreatedAt
	workflowId := ulid.Make()
	workflow.WorkflowId = workflowId
	err := workflowService.workflowRepository.SaveWorkflow(workflow)
	if err != nil {
		return err
	}
	log.Printf("The workflow with id: %s was created", workflowId)
	return nil
}
