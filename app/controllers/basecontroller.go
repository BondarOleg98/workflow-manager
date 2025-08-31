package controllers

import (
	"fmt"
	"log"
	"net/http"
)

type AppController struct {
	workflowController WorkflowController
}

func (appController AppController) InitAppControllers() {
	log.Println("Init the app controllers")
	appController.workflowController.InitWorkflowController()
	http.HandleFunc("/", notFoundHandler)
}

func notFoundHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	responseWriter.WriteHeader(http.StatusNotFound)
	const defaultNotFoundMessage string = "The api request was not found"
	_, err := fmt.Fprint(responseWriter, defaultNotFoundMessage)
	if err != nil {
		log.Printf("%s", err.Error())
	}
}
