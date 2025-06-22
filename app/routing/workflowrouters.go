package routing

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"workflowmanager/app/services"
)

type WorkflowEndpoints struct {
}

func getWorkflowsController(
	responseWriter http.ResponseWriter, _ *http.Request) {
	workflows, err := services.GetWorkflows()
	if err != nil {
		responseWriter.WriteHeader(500)
	}
	encodedWorkflows, _ := json.Marshal(workflows)
	_, err = io.Writer.Write(responseWriter, encodedWorkflows)
	if err != nil {
		responseWriter.WriteHeader(500)
	}
}

func getWorkflowByIdController(
	responseWriter http.ResponseWriter, request *http.Request) {
	workflowUrl := strings.TrimPrefix(request.URL.Path, "/workflows/")
	workflowId := strings.Split(workflowUrl, "/")[3]

	workflow, err := services.GetWorkflowById(workflowId)
	if err != nil {
		responseWriter.WriteHeader(500)
	}
	encodedWorkflow, _ := json.Marshal(workflow)
	_, err = io.Writer.Write(responseWriter, encodedWorkflow)
	if err != nil {
		responseWriter.WriteHeader(500)
	}
}

func (workflowEndpoints WorkflowEndpoints) BaseController() {
	http.HandleFunc("/api/workflows", getWorkflowsController)
	http.HandleFunc("/api/workflows/", getWorkflowByIdController)
}
