package controllers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"workflowmanager/app/components/security"
	"workflowmanager/app/components/services"
	"workflowmanager/app/models"
	"workflowmanager/app/models/filtrartion"
)

type WorkflowController struct {
	workflowService *services.WorkflowService
	preAuthorize    *security.PreAuthorize
}

func NewWorkflowController(
	workflowService *services.WorkflowService,
	preAuthorize *security.PreAuthorize) *WorkflowController {
	return &WorkflowController{
		workflowService: workflowService,
		preAuthorize:    preAuthorize,
	}
}

func (workflowController *WorkflowController) AddWorkflowHandlers() {
	slog.Debug("Add the workflow controller")
	baseWorkflowRoute := "/api/v1/workflows"
	http.Handle(fmt.Sprintf("GET %s", baseWorkflowRoute),
		workflowController.preAuthorize.SecurityFilterChain(http.HandlerFunc(workflowController.getWorkflowsByPagination)))
	http.Handle(fmt.Sprintf("GET %s/{workflowId}", baseWorkflowRoute),
		workflowController.preAuthorize.SecurityFilterChain(http.HandlerFunc(workflowController.getWorkflowById)))
	http.Handle(fmt.Sprintf("DELETE %s/remove/{workflowId}", baseWorkflowRoute),
		workflowController.preAuthorize.SecurityFilterChain(http.HandlerFunc(workflowController.removeWorkflowById)))
	http.Handle(fmt.Sprintf("POST %s/save", baseWorkflowRoute),
		workflowController.preAuthorize.SecurityFilterChain(http.HandlerFunc(workflowController.saveWorkflow)))
	http.Handle(fmt.Sprintf("POST %s", baseWorkflowRoute),
		workflowController.preAuthorize.SecurityFilterChain(http.HandlerFunc(workflowController.getWorkflowsByFiltration)))
}

func (workflowController *WorkflowController) getWorkflowsByPagination(
	responseWriter http.ResponseWriter, request *http.Request) {
	workflowController.preAuthorize.IsAuthorised(responseWriter, request)
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

func (workflowController *WorkflowController) getWorkflowById(
	responseWriter http.ResponseWriter, request *http.Request) {
	workflowController.preAuthorize.IsAuthorised(responseWriter, request)
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

func (workflowController *WorkflowController) removeWorkflowById(responseWriter http.ResponseWriter, request *http.Request) {
	workflowController.preAuthorize.IsAuthorised(responseWriter, request)
	workflowId := request.PathValue("workflowId")
	err := workflowController.workflowService.RemoveWorkflowById(workflowId)
	if err != nil {
		errorMsg := fmt.Sprintf("the workflow with id - %s", workflowId)
		http.Error(responseWriter, errorMsg, http.StatusNotFound)
	} else {
		responseWriter.WriteHeader(http.StatusAccepted)
	}
}

func (workflowController *WorkflowController) saveWorkflow(responseWriter http.ResponseWriter, request *http.Request) {
	workflowController.preAuthorize.IsAuthorised(responseWriter, request)
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

func (workflowController *WorkflowController) getWorkflowsByFiltration(
	responseWriter http.ResponseWriter, request *http.Request) {
	workflowController.preAuthorize.IsAuthorised(responseWriter, request)
	var filtration filtrartion.Filtration
	err := json.NewDecoder(request.Body).Decode(&filtration)
	if err != nil {
		http.Error(responseWriter, "the invalid request payload", http.StatusBadRequest)
		return
	}
	workflows, err := workflowController.workflowService.GetWorkflowsByFiltration(filtration)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	} else {
		responseWriter.WriteHeader(http.StatusOK)
		buildResponseBody(workflows, responseWriter)
	}
}
