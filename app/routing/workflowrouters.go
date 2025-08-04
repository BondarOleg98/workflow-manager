package routing

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"workflowmanager/app/models"
	"workflowmanager/app/services"
)

func getWorkflowsByPaginationController(
	responseWriter http.ResponseWriter, request *http.Request) {
	pageSize, err := parseRequestIntParam(request.URL.Query().Get("page_size"))
	cursor := request.URL.Query().Get("cursor")
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		workflows, err := services.GetWorkflowsByPagination(cursor, pageSize)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
		} else {
			responseHandler(workflows, responseWriter)
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
		responseWriter.WriteHeader(http.StatusOK)
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

func InitBaseController() {
	log.Println("Init base app's controllers")
	http.HandleFunc("GET /api/workflows", getWorkflowsByPaginationController)
	http.HandleFunc("GET /api/workflows/{workflowId}", getWorkflowByIdController)
	http.HandleFunc("DELETE /api/workflows/remove/{workflowId}", removeWorkflowByIdController)
	http.HandleFunc("POST /api/workflow/save", saveWorkflow)
	http.HandleFunc("/", notFoundHandler)
}
