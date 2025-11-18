package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	http.Handle(fmt.Sprintf("GET %s", baseWorkflowRoute),
		taskController.preAuthorize.SecurityFilterChain(http.HandlerFunc(taskController.getTaskByPagination)))
	http.Handle(fmt.Sprintf("GET %s/{taskId}", baseWorkflowRoute),
		taskController.preAuthorize.SecurityFilterChain(http.HandlerFunc(taskController.getTaskById)))
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

func (taskController *TaskController) getTaskByPagination(
	responseWriter http.ResponseWriter, request *http.Request) {
	taskController.preAuthorize.IsAuthorised(responseWriter, request)
	pageSize, err := strconv.Atoi(request.URL.Query().Get("page_size"))
	if err != nil {
		http.Error(responseWriter, "the issue with request param", http.StatusBadRequest)
	} else {
		cursor := request.URL.Query().Get("cursor")
		tasks, err := taskController.taskService.GetTasksByPagination(cursor, pageSize)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
		} else {
			responseWriter.WriteHeader(http.StatusOK)
			buildResponseBody(tasks, responseWriter)
		}
	}
}

func (taskController *TaskController) getTaskById(
	responseWriter http.ResponseWriter, request *http.Request) {
	taskController.preAuthorize.IsAuthorised(responseWriter, request)
	taskId := request.PathValue("taskId")
	task, err := taskController.taskService.GetTaskById(taskId)
	if err != nil {
		errorMsg := fmt.Sprintf("the task with id - %s", taskId)
		http.Error(responseWriter, errorMsg, http.StatusNotFound)
	} else {
		responseWriter.WriteHeader(http.StatusOK)
		buildResponseBody(task, responseWriter)
	}
}
