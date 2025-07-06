package routing

import (
	"encoding/json"
	"io"
	"net/http"
	"workflowmanager/app/models"
	"workflowmanager/app/services"
)

type WorkflowEndpoints struct {
}

func getWorkflowsController(
	responseWriter http.ResponseWriter, _ *http.Request) {
	workflows, err := services.GetWorkflows()
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	} else {
		responseHandler(workflows, responseWriter)
	}
}

func getWorkflowByIdController(
	responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	workflow, err := services.GetWorkflowById(workflowId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
	} else {
		responseHandler(workflow, responseWriter)
	}
}

func removeWorkflowByIdController(responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	err := services.RemoveWorkflowById(workflowId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
	} else {
		responseWriter.WriteHeader(http.StatusAccepted)
	}
}

func saveWorkflow(responseWriter http.ResponseWriter, request *http.Request) {
	requestBody, _ := io.ReadAll(request.Body)
	workflow := models.Workflow{}
	err := json.Unmarshal(requestBody, &workflow)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
	}
	err = services.SaveWorkflow(workflow)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	} else {
		responseWriter.WriteHeader(http.StatusCreated)
	}
}

func (workflowEndpoints WorkflowEndpoints) BaseController() {
	http.HandleFunc("GET /api/workflows", getWorkflowsController)
	http.HandleFunc("GET /api/workflows/{workflowId}", getWorkflowByIdController)
	http.HandleFunc("DELETE /api/workflows/remove/{workflowId}", removeWorkflowByIdController)
	http.HandleFunc("POST /api/workflow/save", saveWorkflow)
	http.HandleFunc("/", notFoundHandler)
}
