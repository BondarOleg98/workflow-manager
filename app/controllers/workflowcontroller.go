package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"workflowmanager/app/models"
	"workflowmanager/app/services"
)

type WorkflowController struct {
	workflowService services.WorkflowService
}

func InitWorkflowController() WorkflowController {
	return WorkflowController{
		workflowService: services.InitWorkflowService(),
	}
}
func (workflowController WorkflowController) AddWorkflowHandlers() {
	log.Println("Add the workflow controller")
	baseWorkflowRoute := "/api/v1/workflows"
	http.HandleFunc(fmt.Sprintf("GET %s", baseWorkflowRoute), workflowController.getWorkflowsByPagination)
	http.HandleFunc(fmt.Sprintf("GET %s/{workflowId}", baseWorkflowRoute), workflowController.getWorkflowById)
	http.HandleFunc(fmt.Sprintf("DELETE %s/remove/{workflowId}", baseWorkflowRoute), workflowController.removeWorkflowById)
	http.HandleFunc(fmt.Sprintf("POST %s/save", baseWorkflowRoute), workflowController.saveWorkflow)
}

func (workflowController WorkflowController) getWorkflowsByPagination(
	responseWriter http.ResponseWriter, request *http.Request) {
	pageSize, err := strconv.Atoi(request.URL.Query().Get("page_size"))
	if err != nil {
		http.Error(responseWriter, "the issue with request param", http.StatusBadRequest)
	} else {
		cursor := request.URL.Query().Get("cursor")
		workflows, err := workflowController.workflowService.GetWorkflowsByPagination(cursor, pageSize)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
		} else {
			responseWriter.WriteHeader(http.StatusOK)
			buildResponseBody(workflows, responseWriter)
		}
	}
}

func (workflowController WorkflowController) getWorkflowById(
	responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	workflow, err := workflowController.workflowService.GetWorkflowById(workflowId)
	if err != nil {
		errorMsg := fmt.Sprintf("the workflow with id - %s", workflowId)
		http.Error(responseWriter, errorMsg, http.StatusNotFound)
	} else {
		responseWriter.WriteHeader(http.StatusOK)
		buildResponseBody(workflow, responseWriter)
	}
}

func (workflowController WorkflowController) removeWorkflowById(responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	err := workflowController.workflowService.RemoveWorkflowById(workflowId)
	if err != nil {
		errorMsg := fmt.Sprintf("the workflow with id - %s", workflowId)
		http.Error(responseWriter, errorMsg, http.StatusNotFound)
	} else {
		responseWriter.WriteHeader(http.StatusAccepted)
	}
}

func (workflowController WorkflowController) saveWorkflow(responseWriter http.ResponseWriter, request *http.Request) {
	var workflow models.Workflow
	err := json.NewDecoder(request.Body).Decode(&workflow)
	if err != nil {
		http.Error(responseWriter, "the invalid request payload", http.StatusBadRequest)
		return
	}
	err = workflowController.workflowService.SaveWorkflow(workflow)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	} else {
		responseWriter.WriteHeader(http.StatusCreated)
	}
}
