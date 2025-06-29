package services

import (
	"log"
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

func RemoveWorkflowById(workflowId string) {
	repository.RemoveWorkflowById(workflowId)
}
