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

type workflowController struct {
}

func (workflowController workflowController) InitWorkflowController() {
	log.Println("Init the workflow controllers")
	baseWorkflowRoute := "/api/v1/workflows"
	http.HandleFunc(fmt.Sprintf("GET %s", baseWorkflowRoute), getWorkflowsByPagination)
	http.HandleFunc(fmt.Sprintf("GET %s/{workflowId}", baseWorkflowRoute), getWorkflowById)
	http.HandleFunc(fmt.Sprintf("DELETE %s/remove/{workflowId}", baseWorkflowRoute), removeWorkflowById)
	http.HandleFunc(fmt.Sprintf("POST %s/save", baseWorkflowRoute), saveWorkflow)
}

func getWorkflowsByPagination(
	responseWriter http.ResponseWriter, request *http.Request) {
	pageSize, err := strconv.Atoi(request.URL.Query().Get("page_size"))
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		cursor := request.URL.Query().Get("cursor")
		workflows, err := services.GetWorkflowsByPagination(cursor, pageSize)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
		} else {
			responseComposer(workflows, responseWriter)
		}
	}
}

func getWorkflowById(
	responseWriter http.ResponseWriter, request *http.Request) {
	workflowId := request.PathValue("workflowId")
	workflow, err := services.GetWorkflowById(workflowId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
	} else {
		responseComposer(workflow, responseWriter)
	}
}

func removeWorkflowById(responseWriter http.ResponseWriter, request *http.Request) {
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
