package services

import (
	"workflowmanager/models"
	"workflowmanager/repository"
)

func GetWorkflows() ([]models.Workflow, error) {
	workflows, err := repository.GetWorkflows()
	return workflows, err
}
