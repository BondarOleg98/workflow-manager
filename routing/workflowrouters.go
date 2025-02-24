package routing

import (
	"io"
	"net/http"
)

type WorkflowEndpoints struct {
}

func getWorkflowController(
	responseWriter http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(responseWriter, "[]\n")
	if err != nil {
		return
	}
}

func (workflowEndpoints WorkflowEndpoints) BaseController() {
	http.HandleFunc("/api/workflows", getWorkflowController)
}
