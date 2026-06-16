package components

import (
	"log/slog"
	"net/http"
	"workflowmanager/app/components/controllers"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/components/security"
	"workflowmanager/app/components/services"
	"workflowmanager/app/persistence/db"
)

func InitAppComponents() {
	slog.Info("Init app components")
	dbInstance := db.GetDatabaseInstance()

	workflowRepository := repository.NewPostgresWorkflowRepository(dbInstance)
	taskRepository := repository.NewPostgresTaskRepository(dbInstance)
	userRepository := repository.NewPostgresUserRepository(dbInstance)
	refreshTokenRepository := repository.NewPostgresRefreshTokenRepository(dbInstance)

	workflowService := services.NewWorkflowService(workflowRepository)
	taskService := services.NewTaskService(taskRepository)
	authService := services.NewAuthService(userRepository, refreshTokenRepository)

	preAuthorize := security.NewPreAuthorize(authService)
	authController := controllers.NewAuthController(authService)
	workflowController := controllers.NewWorkflowController(workflowService, preAuthorize, GetValidatorInstance())
	taskController := controllers.NewTaskController(taskService, preAuthorize)

	authController.AddAuthHandlers()
	workflowController.AddWorkflowHandlers()
	taskController.AddTaskHandlers()

	http.HandleFunc("/", notFoundHandler)
}

func notFoundHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	http.Error(responseWriter, "the api request was not found", http.StatusNotFound)
}
