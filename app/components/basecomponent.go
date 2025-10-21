package components

import (
	"log"
	"net/http"
	"workflowmanager/app/components/controllers"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/components/services"
	"workflowmanager/app/db"
)

func InitAppComponents() {
	log.Println("Init app components")
	dbInstance := db.GetDatabaseInstance()

	workflowRepository := repository.NewWorkflowRepository(dbInstance)
	workflowService := services.NewWorkflowService(workflowRepository)
	workflowController := controllers.NewWorkflowController(workflowService)

	authRepository := repository.NewAuthRepository(dbInstance)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	authController.AddAuthHandlers()
	authChecker := controllers.NewAuthChecker(authService)
	workflowController.AddWorkflowHandlers(authChecker)

	http.HandleFunc("/", notFoundHandler)
}

func notFoundHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	http.Error(responseWriter, "the api request was not found", http.StatusNotFound)
}
