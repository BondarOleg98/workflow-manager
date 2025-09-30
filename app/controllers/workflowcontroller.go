package controllers

import (
	"encoding/json"
	"fmt"
	"io"
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
func (workflowController WorkflowController) AddWorkflowHandlers () {
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
		responseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		cursor := request.URL.Query().Get("cursor")
		workflows, err := workflowController.workflowService.GetWorkflowsByPagination(cursor, pageSize)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
		} else {
			responseComposer(workflows, responseWriter)
		}
	}
}

func (workflowController WorkflowController) getWorkflowById(
	responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	workflow, err := workflowController.workflowService.GetWorkflowById(workflowId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
	} else {
		responseComposer(workflow, responseWriter)
	}
}

func (workflowController WorkflowController) removeWorkflowById(responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	err := workflowController.workflowService.RemoveWorkflowById(workflowId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
	} else {
		responseWriter.WriteHeader(http.StatusAccepted)
	}
}

func (workflowController WorkflowController) saveWorkflow(responseWriter http.ResponseWriter, request *http.Request) {
	requestBody, _ := io.ReadAll(request.Body)
	workflow := models.Workflow{}
	err := json.Unmarshal(requestBody, &workflow)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
	}
	err = workflowController.workflowService.SaveWorkflow(workflow)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	} else {
		responseWriter.WriteHeader(http.StatusCreated)
	}
}
