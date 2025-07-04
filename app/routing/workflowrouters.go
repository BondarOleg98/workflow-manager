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
		encodedWorkflows, _ := json.Marshal(workflows)
		_, err = io.Writer.Write(responseWriter, encodedWorkflows)
		responseWriter.WriteHeader(http.StatusOK)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func getWorkflowByIdController(
	responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	workflow, err := services.GetWorkflowById(workflowId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
	} else {
		encodedWorkflow, _ := json.Marshal(workflow)
		_, err = io.Writer.Write(responseWriter, encodedWorkflow)
		responseWriter.WriteHeader(http.StatusOK)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func removeWorkflowByIdController(responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	_, err := services.RemoveWorkflowById(workflowId)
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
}

func (workflowEndpoints WorkflowEndpoints) BaseController() {
	http.HandleFunc("GET /api/workflows", getWorkflowsController)
	http.HandleFunc("GET /api/workflows/{workflowId}", getWorkflowByIdController)
	http.HandleFunc("DELETE /api/workflows/remove/{workflowId}", removeWorkflowByIdController)
	http.HandleFunc("POST /api/workflow/save", saveWorkflow)
}
