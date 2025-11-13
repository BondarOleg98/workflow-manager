package services

import (
	"log"
	"workflowmanager/app/components/repository"
)

type TaskService struct {
	postgresTaskRepository *repository.PostgresTaskRepository
}

func NewTaskService(postgresTaskRepository *repository.PostgresTaskRepository) *TaskService {
	return &TaskService{
		postgresTaskRepository: postgresTaskRepository,
	}
}

func (taskService *TaskService) RemoveTaskById(taskId string) error {
	rowsTasksAffected, err := taskService.postgresTaskRepository.RemoveTaskById(taskId)
	if err == nil {
		log.Printf("Task by id - %s was removed %d", taskId, rowsTasksAffected)
		return nil
	}
	return err
}
