package controllers

import (
	"fmt"
	"log"
	"net/http"
	"workflowmanager/app/components/security"
	"workflowmanager/app/components/services"
)

type TaskController struct {
	taskService  *services.TaskService
	preAuthorize *security.PreAuthorize
}

func NewTaskController(
	taskService *services.TaskService,
	preAuthorize *security.PreAuthorize) *TaskController {
	return &TaskController{
		taskService:  taskService,
		preAuthorize: preAuthorize,
	}
}

func (taskController *TaskController) AddTaskHandlers() {
	log.Println("Add the task controller")
	baseWorkflowRoute := "/api/v1/task"
	http.Handle(fmt.Sprintf("DELETE %s/remove/{taskId}", baseWorkflowRoute),
		taskController.preAuthorize.SecurityFilterChain(http.HandlerFunc(taskController.removeTaskById)))
}

func (taskController *TaskController) removeTaskById(responseWriter http.ResponseWriter, request *http.Request) {
	taskController.preAuthorize.IsAuthorised(responseWriter, request)
	taskId := request.PathValue("taskId")
	err := taskController.taskService.RemoveTaskById(taskId)
	if err != nil {
		errorMsg := fmt.Sprintf("the task with id - %s", taskId)
		http.Error(responseWriter, errorMsg, http.StatusNotFound)
	} else {
		responseWriter.WriteHeader(http.StatusAccepted)
	}
}
