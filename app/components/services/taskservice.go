package services

import (
	"log"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/models"
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

func (taskService *TaskService) GetTasksByPagination(cursor string, pageSize int) (tasks []models.Task, err error) {
	tasks, err = taskService.postgresTaskRepository.GetTasksByPagination(cursor, pageSize)
	if err == nil {
		log.Printf("Tasks were retrieved")
	}
	return
}

func (taskService *TaskService) GetTaskById(taskId string) (task models.Task, err error) {
	task, err = taskService.postgresTaskRepository.GetTaskById(taskId)
	if err == nil {
		log.Printf("Task by id - %s was retrieved", taskId)
	}
	return
}
