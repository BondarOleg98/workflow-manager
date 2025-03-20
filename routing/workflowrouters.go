package routing

import (
	"encoding/json"
	"io"
	"net/http"
	"workflowmanager/services"
)

type WorkflowEndpoints struct {
}

func getWorkflowController(
	responseWriter http.ResponseWriter, _ *http.Request) {
	workflows, _ := json.Marshal(services.GetWorkflows())
	_, err := io.Writer.Write(responseWriter, workflows)
	if err != nil {
		return
	}
}

func (workflowEndpoints WorkflowEndpoints) BaseController() {
	http.HandleFunc("/api/workflows", getWorkflowController)
}
