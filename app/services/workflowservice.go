package services

import (
	"workflowmanager/app/models"
	"workflowmanager/app/repository"
)

func GetWorkflows() ([]models.Workflow, error) {
	workflows, err := repository.GetWorkflows()
	return workflows, err
}
