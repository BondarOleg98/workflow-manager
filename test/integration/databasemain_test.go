package integration

import (
	"testing"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/components/services"
	"workflowmanager/app/db"
)

func TestDatabaseMain(test *testing.T) {
	pool := db.Pool{
		DriverName:    "postgres",
		ConnectionUrl: "postgres://postgres:postgres@localhost/workflow_manager?sslmode=disable",
	}
	db.InitDatabaseInstance(pool)
	defer db.CloseDatabaseConnection()
	workflowRepository := repository.NewWorkflowRepository(db.GetDatabaseInstance())
	workflowService := services.NewWorkflowService(workflowRepository)
	workflowServiceTest := NewWorkflowServiceTest(workflowService)

	test.Run("TestCrdOperations", func(test *testing.T) {
		workflowServiceTest.testCrdOperations(test)
	})
	test.Run("TestCorrectDatabaseInstance", func(t *testing.T) {
		testCorrectDatabaseInstance(test)
	})
}
