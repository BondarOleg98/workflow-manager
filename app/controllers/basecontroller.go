package controllers

import (
	"log"
	"net/http"
)

type AppController struct {
}

func (appController AppController) InitAppControllers() {
	log.Println("Init the app controllers")
	authController := InitAuthController()
	authController.AddAuthHandlers()
	authService := authController.authService
	InitWorkflowController().AddWorkflowHandlers(authService)
	http.HandleFunc("/", notFoundHandler)
}

func notFoundHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	http.Error(responseWriter, "the api request was not found", http.StatusNotFound)
}
