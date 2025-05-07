package routing

import (
	"encoding/json"
	"io"
	"net/http"
	"workflowmanager/app/services"
)

type WorkflowEndpoints struct {
}

func getWorkflowController(
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

func (workflowEndpoints WorkflowEndpoints) BaseController() {
	http.HandleFunc("/api/workflows", getWorkflowController)
}
