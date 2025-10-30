package components

import (
	"log"
	"net/http"
	"workflowmanager/app/components/controllers"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/components/security"
	"workflowmanager/app/components/services"
	"workflowmanager/app/db"
)

func InitAppComponents() {
	log.Println("Init app components")
	dbInstance := db.GetDatabaseInstance()

	workflowRepository := repository.NewWorkflowRepository(dbInstance)
	userRepository := repository.NewPostgresUserRepository(dbInstance)
	refreshTokenRepository := repository.NewPostgresRefreshTokenRepository(dbInstance)

	workflowService := services.NewWorkflowService(workflowRepository)
	authService := services.NewAuthService(userRepository, refreshTokenRepository)

	preAuthorize := security.NewPreAuthorize(authService)
	authController := controllers.NewAuthController(authService)
	workflowController := controllers.NewWorkflowController(workflowService, preAuthorize)

	authController.AddAuthHandlers()
	workflowController.AddWorkflowHandlers()

	http.HandleFunc("/", notFoundHandler)
}

func notFoundHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	http.Error(responseWriter, "the api request was not found", http.StatusNotFound)
}
