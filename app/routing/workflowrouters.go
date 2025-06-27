package routing

import (
	"encoding/json"
	"io"
	"net/http"
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
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func removeWorkflowByIdController(responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	services.RemoveWorkflowById(workflowId)
}

func (workflowEndpoints WorkflowEndpoints) BaseController() {
	http.HandleFunc("GET /api/workflows", getWorkflowsController)
	http.HandleFunc("GET /api/workflows/{workflowId}", getWorkflowByIdController)
	http.HandleFunc("DELETE /api/workflows/{workflowId}", removeWorkflowByIdController)
}
