package services

import (
	"github.com/oklog/ulid/v2"
	"log"
	"time"
	"workflowmanager/app/models"
	"workflowmanager/app/repository"
)

func GetWorkflows() (workflows []models.Workflow, err error) {
	workflows, err = repository.GetWorkflows()
	if err == nil {
		log.Printf("Workflows were retrieved")
	}
	return
}

func GetWorkflowById(workflowId string) (workflow models.Workflow, err error) {
	workflow, err = repository.GetWorkflowById(workflowId)
	if err == nil {
		log.Printf("Workflow by id - %s was retrieved", workflowId)
	}
	return
}

func RemoveWorkflowById(workflowId string) error {
	rowsWorkflowsAffected, err := repository.RemoveWorkflowById(workflowId)
	if err == nil {
		log.Printf("Workflow by id - %s was removed %d", workflowId, rowsWorkflowsAffected)
		return nil
	}
	return err
}

func SaveWorkflow(workflow models.Workflow) error {
	workflow.CreatedAt = time.Now()
	workflow.UpdatedAt = workflow.CreatedAt
	workflowId := ulid.Make()
	workflow.WorkflowId = workflowId
	err := repository.SaveWorkflow(workflow)
	if err != nil {
		return err
	}
	log.Printf("The workflow with id: %s was created", workflowId)
	return nil
}
