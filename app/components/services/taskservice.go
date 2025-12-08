package services

import (
	"log"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/models"
)

type TaskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

func (taskService *TaskService) RemoveTaskById(taskId string) error {
	rowsTasksAffected, err := taskService.taskRepository.RemoveTaskById(taskId)
	if err == nil {
		log.Printf("Task by id - %s was removed %d", taskId, rowsTasksAffected)
		return nil
	}
	return err
}

func (taskService *TaskService) GetTasksByPagination(cursor string, pageSize int) (tasks []models.Task, err error) {
	tasks, err = taskService.taskRepository.GetTasksByPagination(cursor, pageSize)
	if err == nil {
		log.Printf("Tasks were retrieved")
	}
	return
}

func (taskService *TaskService) GetTaskById(taskId string) (task models.Task, err error) {
	task, err = taskService.taskRepository.GetTaskById(taskId)
	if err == nil {
		log.Printf("Task by id - %s was retrieved", taskId)
	}
	return
}

func (taskService *TaskService) GetTasksByWorkflowId(workflowId string) (tasks []models.Task, err error) {
	tasks, err = taskService.taskRepository.GetTasksByWorkflowId(workflowId)
	if err == nil {
		log.Printf("Tasks by workflowId %s were retrieved", workflowId)
	}
	return
}
